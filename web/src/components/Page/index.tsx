import React from 'react'
import Header from '../Header'
import MainContent from '../MainContent'

interface Props {
  children?: React.ReactNode
}

export default function Page({children}: Props) {
  return (
    <div className="Page"><Header/><MainContent>{children}</MainContent></div>
  )
}
