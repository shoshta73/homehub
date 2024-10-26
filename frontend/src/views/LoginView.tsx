import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useNavigate } from "react-router-dom";

const loginSchema = z.object({
  username: z.string().max(50).optional(),
  email: z.string().email().optional(),
  password: z.string().min(8),
});

function RegisterView() {
  const navigate = useNavigate();

  const form = useForm<z.infer<typeof loginSchema>>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      username: "",
      email: "",
      password: "",
    },
  });

  function onSubmit(values: z.infer<typeof loginSchema>) {
    fetch(`/auth/login`, {
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
    <div className="flex flex-col flex-grow items-center justify-center">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8 w-96" data-testid="login-form">
          <FormField
            control={form.control}
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
            control={form.control}
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
            control={form.control}
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
    </div>
  );
}

export default RegisterView;
