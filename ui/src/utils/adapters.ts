import axios from "axios";
import { matchRoutes } from "react-router-dom";
import z from "zod";

import { ROUTES } from "src/routes";
import { StrictSchema } from "src/utils/types";

export interface APIError {
  message: string;
  status?: number;
}

interface APIErrorResponse {
  error: APIError;
  isSuccessful: false;
}

interface APISuccessfulResponse<D> {
  data: D;
  isSuccessful: true;
}

export type APIResponse<D> = APISuccessfulResponse<D> | APIErrorResponse;

export enum HTTPStatusError {
  Aborted = 0,
  Unauthorized = 401,
  NotFound = 404,
}

export enum HTTPStatusSuccess {
  Accepted = 202,
  Created = 201,
  NoContent = 204,
  OK = 200,
}

export interface ResultOK<D> {
  data: D;
  status: HTTPStatusSuccess.OK;
}

export interface ResultCreated<D> {
  data: D;
  status: HTTPStatusSuccess.Created;
}

interface ResultAccepted<D> {
  data: D;
  status: HTTPStatusSuccess.Accepted;
}

export const resultAcceptedNull = StrictSchema<ResultAccepted<null>>()(
  z.object({
    data: z.null(),
    status: z.literal(HTTPStatusSuccess.Accepted),
  })
);

interface ResultNoContent<D> {
  data: D;
  status: HTTPStatusSuccess.NoContent;
}

export const resultNoContent = StrictSchema<ResultNoContent<"">>()(
  z.object({
    data: z.literal(""),
    status: z.literal(HTTPStatusSuccess.NoContent),
  })
);

export interface ResponseError {
  data: { message: string };
  status: number;
}

export const responseError = StrictSchema<ResponseError>()(
  z.object({
    data: z.object({ message: z.string() }),
    status: z.number(),
  })
);

export function buildAPIError(error: unknown): APIError {
  if (axios.isCancel(error)) {
    return { message: error.toString(), status: HTTPStatusError.Aborted };
  }

  if (axios.isAxiosError(error)) {
    try {
      // This is a Polygon ID API error.
      const { data, status } = responseError.parse(error.response);
      const { message } = data;

      if (status === HTTPStatusError.Unauthorized) {
        const isAuthorizedPath = Object.values(ROUTES)
          .filter(({ path }) => ROUTES.notFound.path !== path && ROUTES.signIn.path !== path)
          .find(({ path }) => {
            const currentRoute = matchRoutes([{ path }], location.pathname)?.[0];

            return currentRoute;
          });

        if (isAuthorizedPath) {
          window.location.href = ROUTES.signIn.path;
        }
      }

      return { message, status };
    } catch (e) {
      // This catches a CORS or other network error.
      return { message: error.message };
    }
  }

  if (error instanceof Error) {
    // This is an application-level error.
    return { message: error.toString() };
  }

  // This shouldn't happen (catch-all).
  console.error(error);
  return { message: "Unknown error" };
}

export function processZodError<T>(error: z.ZodError<T>, init: string[] = []): string[] {
  return error.errors.reduce((mainAcc, issue): string[] => {
    switch (issue.code) {
      case "invalid_union": {
        return [
          ...mainAcc,
          ...issue.unionErrors.reduce(
            (innerAcc: string[], current: z.ZodError<T>): string[] => [
              ...innerAcc,
              ...processZodError(current, mainAcc),
            ],
            []
          ),
        ];
      }
      default: {
        const errorMsg = issue.path.length
          ? `${issue.message} at ${issue.path.join(".")}`
          : issue.message;
        return [...mainAcc, errorMsg];
      }
    }
  }, init);
}

export const numericBoolean = StrictSchema<0 | 1, boolean>()(
  z.union([z.literal(0), z.literal(1)]).transform((value) => value === 1)
);

export const stringBoolean = StrictSchema<string, boolean>()(
  z.string().transform((value, context) => {
    switch (value) {
      case "true": {
        return true;
      }
      case "false": {
        return false;
      }
      default: {
        context.addIssue({
          code: z.ZodIssueCode.custom,
          fatal: true,
          message: "The provided string input can't be parsed as a boolean",
        });
        return z.NEVER;
      }
    }
  })
);