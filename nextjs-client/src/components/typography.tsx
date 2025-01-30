type TypographyProps = {
    message : string
}

export function TypographyH3({message}:TypographyProps) {
    return (
      <h3 className="scroll-m-20 text-2xl font-semibold tracking-tight">
        {message}
      </h3>
    )
  }

export function TypographyH4({message}:TypographyProps) {
    return (
      <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">
        {message}
      </h4>
    )
}
export function TypographyP({message}:TypographyProps) {
    return (
      <p className="leading-7 [&:not(:first-child)]:mt-6">
        {message}
      </p>
    )
  }
  
  