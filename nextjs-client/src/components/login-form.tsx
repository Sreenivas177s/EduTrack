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
import { useSearchParams } from "next/navigation"
import { UIAlert } from "./common"

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

  const onSubmit = (values:z.infer<typeof SignInSchema>) => {
    signIn("credentials", {redirectTo : "/ui/home",...values})
  }
  const errorMsg = useSearchParams().get("code") === "credentials" ? "Invalid credentials" : null
  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card>
        <CardHeader className="text-center">
          <CardTitle className="text-xl">Login In</CardTitle>
        </CardHeader>
        <CardContent>
        {errorMsg && (
        <UIAlert
          title="Invalid Credentials"
          description="The email or password you entered is incorrect."
          variant="destructive"
        />
      )}
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <div className="grid gap-6">
              <div className="grid gap-6">
                <div className="grid gap-2">
                  {getFormField(form.control, "email_id", "Email","Enter your email")}
                  {getFormField(form.control, "password","Password","Enter your Password")}
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

function getFormField(control: Control<z.infer<typeof SignInSchema>>, name: keyof z.infer<typeof SignInSchema>,displayName:string, desc: string) {
  return (
    <FormField control={control} name={name}
      render={({field}) => {
        return (
          <>
            <FormItem>
                <FormLabel>{displayName}</FormLabel>
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