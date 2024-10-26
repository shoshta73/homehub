import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useNavigate } from "react-router-dom";

const registerSchema = z.object({
  username: z.string().min(3).max(50),
  name: z.string().min(3).max(50),
  email: z.string().email(),
  password: z.string().min(8),
  confirmPassword: z.string().min(8),
});

function RegisterView() {
  const navigate = useNavigate();

  const form = useForm<z.infer<typeof registerSchema>>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      username: "",
      name: "",
      email: "",
      password: "",
      confirmPassword: "",
    },
  });

  function onSubmit(values: z.infer<typeof registerSchema>) {
    if (values.confirmPassword.length !== values.password.length) {
      if (values.confirmPassword.length < values.password.length) {
        form.setError("confirmPassword", {
          message: "ConfirmPassword field is shorter than the password.",
        });
      }
      if (values.confirmPassword.length > values.password.length) {
        form.setError("confirmPassword", {
          message: "ConfirmPassword field is longer than the password.",
        });
      }
      return;
    }

    if (values.password !== values.confirmPassword) {
      form.setError("confirmPassword", {
        message: "Passwords do not match.",
      });
      form.setError("password", {
        message: "Passwords do not match.",
      });
      return;
    }

    fetch(`/auth/register`, {
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
    <div className="flex flex-col items-center justify-center h-full w-full">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8 w-96" data-testid="register-form">
          <FormField
            control={form.control}
            name="username"
            render={({ field }) => (
              <FormItem data-testid="username-field">
                <FormLabel>Username</FormLabel>
                <FormControl>
                  <Input placeholder="username" {...field} />
                </FormControl>
                <FormDescription>This is your username name.</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem data-testid="name-field">
                <FormLabel>Name</FormLabel>
                <FormControl>
                  <Input placeholder="name" {...field} />
                </FormControl>
                <FormDescription>
                  This is your name. It's optional, but if you set this it will be used instead of your username.
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem data-testid="email-field">
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input placeholder="email" {...field} />
                </FormControl>
                <FormDescription>This is your email address.</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem data-testid="password-field">
                <FormLabel>Password</FormLabel>
                <FormControl>
                  <Input type="password" placeholder="password" {...field} />
                </FormControl>
                <FormDescription>This is your password.</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="confirmPassword"
            render={({ field }) => (
              <FormItem data-testid="confirm-password-field">
                <FormLabel>Confirm Password</FormLabel>
                <FormControl>
                  <Input type="password" placeholder="confirm password" {...field} />
                </FormControl>
                <FormDescription>You must type the same password as above.</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <div>
            <Label>Already have an account?</Label>
            <Button type="button" className="w-full" onClick={() => navigate("/login")} data-testid="login-button">
              Login
            </Button>
          </div>

          <Button type="submit" className="w-full" data-testid="register-button">
            Register
          </Button>
        </form>
      </Form>
    </div>
  );
}

export default RegisterView;
