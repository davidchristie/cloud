import React from 'react'

interface Props {
  children?: React.ReactNode
}

export default function PageHeading({ children }: Props) {
  return (
    <h1 className="PageHeading">{children}</h1>
  )
}
