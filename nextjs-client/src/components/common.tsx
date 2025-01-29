import { AlertCircle } from "lucide-react"

import {
  Alert,
  AlertDescription,
  AlertTitle,
} from "@/components/ui/alert"

import { FC } from "react";

interface UIAlertProps {
  title: string;
  description: string;
  variant: "default" | "destructive" | null | undefined;
}

export const UIAlert: FC<UIAlertProps> = ({ title, description, variant }) => {
  return (
    <Alert variant={variant}>
      <AlertCircle className="h-4 w-4" />
      <AlertTitle>{title}</AlertTitle>
      <AlertDescription>
        {description}
      </AlertDescription>
    </Alert>
  )
}
