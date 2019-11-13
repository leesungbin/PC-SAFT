import React from 'react';

type MenuProps = {
  open: boolean,
  children: React.ReactElement[],
}


export default function Menu(props: MenuProps) {
  return (
    <div style={{ ...styles.container, height: props.open ? '100%' : 0 }}>
      {
        props.open ?
          <div style={styles.menuList}>
            {props.children}
          </div> : null
      }
    </div>
  )
}

const styles: { [key: string]: React.CSSProperties } = {
  container: {
    position: 'absolute',
    top: 0,
    left: 0,
    width: '100vw',
    display: 'flex',
    flexDirection: 'column',
    background: 'black',
    opacity: 0.9,
    color: '#fafafa',
    transition: 'height 0.3s ease',
    zIndex: -1,
  },
  menuList: {
    paddingTop: '3rem',
  }
}