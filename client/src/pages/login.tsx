import LoginForm from "@/components/custom/login-form"
import { TypographyH1, TypographyLead, TypographyMuted } from "@/components/custom/typography"
import { ThemeToggle } from "@/components/theme/theme-toggle"


export default function LoginPage() {
  return (
    <div className="flex w-full h-full">
        <ApplicationTitle />
        <LoginSegment />
    </div>
  )
}

function ApplicationTitle() {
    return (
        <div className="bg-secondary basis-3/5">
            <div className="flex flex-col absolute left-10 top-20">
                <TypographyH1 data={import.meta.env.VITE_APP_NAME} extraClasses="text-primary text-6xl" />
                <TypographyLead data={import.meta.env.VITE_APP_DESCRIPTION} extraClasses="mt-2" />
            </div>
            <div className="absolute left-10 bottom-10">
            <TypographyMuted data={`@ ${import.meta.env.VITE_APP_YEAR} ${import.meta.env.VITE_APP_AUTHOR_NAME}`} />
            </div>
        </div>
    )
}

function LoginSegment() {
    return (
        <div className="flex items-center justify-center basis-2/5 h-screen">
            <LoginForm />
            <div className="absolute top-10 right-10">
                <ThemeToggle />
            </div>
        </div>
    );
}
