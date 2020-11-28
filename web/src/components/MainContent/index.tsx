import React from 'react'
import './index.css'

interface Props {
  children?: React.ReactNode
}

export default function MainContent({children}: Props) {
  return (
    <div className="MainContent">{children}</div>
  )
}
