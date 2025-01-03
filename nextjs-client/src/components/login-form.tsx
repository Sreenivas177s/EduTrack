"use client"

import { cn, SignInSchema } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Control, useForm } from "react-hook-form"
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "./ui/form"
import { signIn } from "next-auth/react"

export function LoginForm({
  className,
  ...props
}: React.ComponentPropsWithoutRef<"div">) {

  const form =useForm<z.infer<typeof SignInSchema>>({
    resolver: zodResolver(SignInSchema),
    defaultValues: {
      email_id: "",
      password: "",
    },
  })

  const credentialsAction = (values: FormData) => {
    signIn("credentials", {redirectTo : "/home"}, values)
  }
  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card>
        <CardHeader className="text-center">
          <CardTitle className="text-xl">Login In</CardTitle>
        </CardHeader>
        <CardContent>
        <Form {...form}>
          <form action={credentialsAction}>
            <div className="grid gap-6">
              <div className="grid gap-6">
                <div className="grid gap-2">
                  {getFormField(form.control, "email_id", "Enter your email")}
                  {getFormField(form.control, "password", "Enter your Password")}
                </div>
                <Button type="submit" className="w-full">
                  Login
                </Button>
              </div>
              <div className="text-center text-sm">
                Don&apos;t have an account?{" "}
                <a href="#" className="underline underline-offset-4">
                  Sign up
                </a>
              </div>
            </div>
          </form>
          </Form>
        </CardContent>
      </Card>
      <div className="text-balance text-center text-xs text-muted-foreground [&_a]:underline [&_a]:underline-offset-4 [&_a]:hover:text-primary  ">
        By clicking continue, you agree to our <a href="#">Terms of Service</a>{" "}
        and <a href="#">Privacy Policy</a>.
      </div>
    </div>
  )
}

function getFormField(control: Control<z.infer<typeof SignInSchema>>, name: keyof z.infer<typeof SignInSchema>, desc: string) {
  return (
    <FormField control={control} name={name}
      render={({field}) => {
        return (
          <>
            <FormItem>
                <FormLabel>{name}</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>
                <FormDescription>
                  {desc}
                </FormDescription>
                <FormMessage />
              </FormItem>

          </>
        );
      }} />
  );
}