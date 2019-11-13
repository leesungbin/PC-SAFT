import React, { useState } from 'react';
import { NavLink } from 'react-router-dom';

type MenuItemProps = {
  delay: string,
  value: { title: string, to: string, exact: boolean}
  onClick: () => void,
}

export default function MenuItem(props: MenuItemProps) {
  const [hover, setHover] = useState(false);
  const handleHover = () => {
    setHover(!hover);
  }
  const styles: {[key: string]: React.CSSProperties} = {
    container: {
      opacity: 0,
      animation: '1s appear forwards',
      animationDelay: props.delay,
      height: 70,
      zIndex: 100,
    },
    menuItem: {
      fontFamily: `'Open Sans', sans-serif`,
      fontSize: '1.2rem',
      padding: '1rem 0',
      margin: '0 5%',
      cursor: 'pointer',
      color: hover ? 'gray' : '#fafafa',
      transition: 'color 0.2s ease-in-out',
      animation: '0.5s slideIn forwards',
      animationDelay: props.delay,
    },
    line: {
      width: '90%',
      height: '1px',
      background: 'gray',
      margin: '0 auto',
      animation: '0.5s shrink forwards',
      animationDelay: props.delay,
    }
  }
  return (
    <div style={styles.container}>
      <NavLink
        exact={props.value.exact}
        to={props.value.to}
        style={styles.menuItem}
        activeStyle={{fontWeight: 'bold'}}
        onMouseEnter={() => { handleHover(); }}
        onMouseLeave={() => { handleHover(); }}
        onClick={() => props.onClick()}
      >
        {props.value.title}
      </NavLink>
      <div style={styles.line} />
    </div>
  )

}