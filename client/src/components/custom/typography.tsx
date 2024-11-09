

interface TypographyProps {
    data: string
    dataArray?: string[]
    extraClasses?: string
}

export function TypographyH1({ data, extraClasses }: TypographyProps) {
    return (
      <h1 className={`scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl ${extraClasses}`}>
        {data}
      </h1>
    )
  }

export function TypographyH2({ data, extraClasses }: TypographyProps) {
    return (
        <h2 className={`scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight first:mt-0 ${extraClasses}`}>
            {data}
        </h2>
    )
  }

  export function TypographyH3({ data, extraClasses }: TypographyProps  ) {
    return (
      <h3 className={`scroll-m-20 text-2xl font-semibold tracking-tight ${extraClasses}`}>
        {data}
      </h3>
    )
  }

  export function TypographyH4({ data, extraClasses }: TypographyProps  ) {
    return (
      <h4 className={`scroll-m-20 text-xl font-semibold tracking-tight ${extraClasses}`}>
        {data}
      </h4>
    )
  }
  export function TypographyP({ data, extraClasses }: TypographyProps) {
    return (
      <p className={`leading-7 [&:not(:first-child)]:mt-6 ${extraClasses}`}>
        {data}
      </p>
    )
  }
  export function TypographyBlockquote({ data, extraClasses }: TypographyProps) {
    return (
      <blockquote className={`mt-6 border-l-2 pl-6 italic ${extraClasses}`}>
        {data}
      </blockquote>
    )
  }
  
  export function TypographyList({ dataArray, extraClasses }: TypographyProps) {
    return (
      <ul className={`my-6 ml-6 list-disc [&>li]:mt-2 ${extraClasses}`}>
        {dataArray?.map((item, index) => (
          <li key={index}>{item}</li>
        ))}
      </ul>
    )
  }
  export function TypographyInlineCode({ data, extraClasses }: TypographyProps  ) {
    return (
      <code className={`relative rounded bg-muted px-[0.3rem] py-[0.2rem] font-mono text-sm font-semibold ${extraClasses}`}>
        {data}
      </code>
    )
  }
  export function TypographyLead({ data, extraClasses }: TypographyProps    ) {
    return (
      <p className={`text-xl text-muted-foreground ${extraClasses}`}>
        {data}
      </p>
    )
  }
  export function TypographyLarge({ data, extraClasses }: TypographyProps) {
    return <div className={`text-lg font-semibold ${extraClasses}`}>{data}</div>
  }
  export function TypographySmall({ data, extraClasses }: TypographyProps) {
    return (
      <small className={`text-sm font-medium leading-none ${extraClasses}`}>
        {data}
      </small>
    )
  }
  export function TypographyMuted({ data, extraClasses }: TypographyProps   ) {
    return (
      <p className={`text-sm text-muted-foreground ${extraClasses}`}>
        {data}
      </p>
    )
  }
  