import React, { useState } from 'react';
import { NavLink, Link } from 'react-router-dom';
import MenuButton from './MenuButton';
import Menu from './Menu';
import MenuItem from './MenuItem';


function Header({ width }: { width: number }) {
  const [open, setOpen] = useState(false);

  const menu = [{
    title: 'Program',
    to: '/',
    exact: true,
  }, {
    title: 'Documentation',
    to: '/docs',
    exact: false,
  }, {
    title: 'Database',
    to: 'db',
    exact: false,
  }];

  // for mobile
  const menuItems = menu.map((val, index) => {
    return (
      <MenuItem
        key={index}
        delay={`${index * 0.1}s`}
        value={val}
        onClick={() => setOpen(false)}
      />)
  });
  return (
    <div style={style.root}>
      <div style={style.logo}><Link to="/" style={style.logolink}>SAFT-GO</Link></div>
      {width > 850 ?
        <ul style={style.linkGroup}>
          {menu.map((e, i) => (
            <li key={i}><NavLink exact={e.exact} to={e.to} style={style.link} activeStyle={{ fontWeight: 'bold' }}>{e.title}</NavLink></li>
          ))}
        </ul>
        :
        <div style={{zIndex: 99}}>
          <Menu open={open}>
            {menuItems}
          </Menu>
          <MenuButton open={open} onClick={() => setOpen(!open)} color={open ? 'white': 'black'}/>
        </div>
      }
    </div>
  );
}
export default Header;

const style: { [key: string]: React.CSSProperties } = {
  root: {
    backgroundColor: 'rgba(174,196,255,0.22)',
    height: 76,
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingLeft: '8%',
    paddingRight: '8%',
  },
  logo: {
    display: 'flex',
    fontSize: 30,
    alignItems: 'center'
  },
  logolink: {
    color: 'black',
    textDecoration: 'none',
  },
  linkGroup: {
    display: 'flex',
    width: 500,
    alignItems: 'center',
    justifyContent: 'space-between',
    listStyleType: 'none',
  },
  link: {
    display: 'flex',
    fontSize: 26,
    color: 'black',
    textDecoration: 'none',
    // marginRight: 30
  }
};