import React from 'react';
import { NavLink, Link } from 'react-router-dom';
import MenuButton from './MenuButton';

function Header({ width }: { width: number }) {
  return (
    <div style={style.root}>
      <div style={style.logo}><Link to="/" style={style.logolink}>SAFT-GO</Link></div>
      {width > 850 ?
        <ul style={style.linkGroup}>
          <li><NavLink exact to="/" style={style.link} activeStyle={{ fontWeight: 'bold' }}>Program</NavLink></li>
          <li><NavLink to="/docs" style={style.link} activeStyle={{ fontWeight: 'bold' }}>Documentation</NavLink></li>
          <li><NavLink to="/db" style={style.link} activeStyle={{ fontWeight: 'bold' }}>Database</NavLink></li>
        </ul>
        :
        <MenuButton open={false} />
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
    paddingLeft: 97,
    paddingRight: 67,
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