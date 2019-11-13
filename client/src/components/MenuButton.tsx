import React, { CSSProperties } from 'react';

type MenuButtonProps = {
  open: boolean,
  color?: string,
  onClick?: () => void,
};

export default function MenuButton(props: MenuButtonProps) {
  const color = props.color? props.color : 'black';
  return (
    <div style={style.container}
      onClick={props.onClick ? props.onClick :
        () => { console.log('hi') }}>
      <div style={{
        ...style.line, background: color, ...style.lineTop,
        transform: props.open ? 'rotate(45deg)' : 'none',
      }} />
      <div style={{
        ...style.line, background: color, ...style.lineMiddle,
        opacity: props.open ? 0 : 1,
        transform: props.open ? 'translateX(-16px)' : 'none',
      }} />
      <div style={{
        ...style.line, background: color, ...style.lineBottom,
        transform: props.open ? 'translateX(-1px) rotate(-45deg)' : 'none',
      }} />
    </div>
  )
}

const style: { [key: string]: CSSProperties } = {
  container: {
    height: '32px',
    width: '32px',
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    cursor: 'pointer',
    padding: '4px',
    zIndex: 100,
  },
  line: {
    height: '2px',
    width: '30px',
    transition: 'all 0.2s ease',
  },
  lineTop: {
    transformOrigin: 'top left',
    marginBottom: '8px',
  },
  lineBottom: {
    transformOrigin: 'top left',
    marginTop: '8px',
  },
}