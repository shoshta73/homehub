import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import { Pagination, PaginationContent, PaginationItem, PaginationLink } from "@/components/ui/pagination";

const usernameSchema = z.object({
  username: z.string().max(50),
  password: z.string().min(8),
});

const emailSchema = z.object({
  email: z.string().email(),
  password: z.string().min(8),
});

function RegisterView() {
  const [method, setMethod] = useState<"username" | "email">("username");

  const navigate = useNavigate();

  const usernameForm = useForm<z.infer<typeof usernameSchema>>({
    resolver: zodResolver(usernameSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  const emailForm = useForm<z.infer<typeof emailSchema>>({
    resolver: zodResolver(emailSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  function emailOnSubmit(values: z.infer<typeof emailSchema>) {
    fetch(`${import.meta.env.VITE_API_URL ?? ""}/auth/login/email`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(values),
    })
      .then((res) => res.text())
      .then((text) => {
        if (text == "OK") {
          navigate("/home");
        }
      })
      .catch((err) => console.error(err));
  }

  function usernameOnSubmit(values: z.infer<typeof usernameSchema>) {
    fetch(`${import.meta.env.VITE_API_URL ?? ""}/auth/login/username`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(values),
    })
      .then((res) => res.text())
      .then((text) => {
        if (text == "OK") {
          navigate("/home");
        }
      })
      .catch((err) => console.error(err));
  }

  return (
    <>
      <div className="flex flex-col flex-grow items-center justify-center">
        <Pagination>
          <PaginationContent>
            <PaginationItem>
              <PaginationLink
                className="w-fit p-1"
                onClick={() => setMethod("username")}
                isActive={method === "username"}
              >
                Username
              </PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationLink className="w-fit p-1" onClick={() => setMethod("email")} isActive={method === "email"}>
                Email
              </PaginationLink>
            </PaginationItem>
          </PaginationContent>
        </Pagination>
        {method === "username" && (
          <Form {...usernameForm}>
            <form
              onSubmit={usernameForm.handleSubmit(usernameOnSubmit)}
              className="space-y-8 w-96"
              data-testid="login-form"
            >
              <FormField
                control={usernameForm.control}
                name="username"
                render={({ field }) => (
                  <FormItem data-testid="username-field">
                    <FormLabel data-testid="username-label">Username</FormLabel>
                    <FormControl>
                      <Input placeholder="username" {...field} data-testid="username-input" />
                    </FormControl>
                    <FormDescription data-testid="username-description">This is your username name.</FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={usernameForm.control}
                name="password"
                render={({ field }) => (
                  <FormItem data-testid="password-field">
                    <FormLabel data-testid="password-label">Password</FormLabel>
                    <FormControl>
                      <Input type="password" placeholder="password" {...field} data-testid="password-input" />
                    </FormControl>
                    <FormDescription data-testid="password-description">This is your password.</FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <div>
                <Label>Dont have an account?</Label>
                <Button
                  type="button"
                  className="w-full"
                  onClick={() => navigate("/register")}
                  data-testid="register-button"
                >
                  Register
                </Button>
              </div>

              <Button type="submit" className="w-full" data-testid="submit-button">
                Login
              </Button>
            </form>
          </Form>
        )}
        {method === "email" && (
          <Form {...emailForm}>
            <form onSubmit={emailForm.handleSubmit(emailOnSubmit)} className="space-y-8 w-96" data-testid="login-form">
              <FormField
                control={emailForm.control}
                name="email"
                render={({ field }) => (
                  <FormItem data-testid="email-field">
                    <FormLabel data-testid="email-label">Email</FormLabel>
                    <FormControl>
                      <Input placeholder="email" {...field} data-testid="email-input" />
                    </FormControl>
                    <FormDescription data-testid="email-description">This is your email address.</FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={emailForm.control}
                name="password"
                render={({ field }) => (
                  <FormItem data-testid="password-field">
                    <FormLabel data-testid="password-label">Password</FormLabel>
                    <FormControl>
                      <Input type="password" placeholder="password" {...field} data-testid="password-input" />
                    </FormControl>
                    <FormDescription data-testid="password-description">This is your password.</FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <div>
                <Label>Dont have an account?</Label>
                <Button
                  type="button"
                  className="w-full"
                  onClick={() => navigate("/register")}
                  data-testid="register-button"
                >
                  Register
                </Button>
              </div>

              <Button type="submit" className="w-full" data-testid="submit-button">
                Login
              </Button>
            </form>
          </Form>
        )}
      </div>
    </>
  );
}

export default RegisterView;
