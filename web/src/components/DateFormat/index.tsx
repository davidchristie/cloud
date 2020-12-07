import React from 'react'

interface Props {
  value: string
}

export default function DateFormat({ value }: Props) {
  return <>{new Date(value).toLocaleString()}</>
}
