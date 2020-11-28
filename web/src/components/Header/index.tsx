import React from 'react'
import { NavLink } from 'react-router-dom'
import { getHomePageUrl, getProductListPageUrl } from '../../utilities'
import './index.css'

export default function Header() {
  return (
    <div className="Header">
      <NavLink exact to={getHomePageUrl()}>Home</NavLink>
      <NavLink exact to={getProductListPageUrl()}>Products</NavLink>
    </div>
  )
}
