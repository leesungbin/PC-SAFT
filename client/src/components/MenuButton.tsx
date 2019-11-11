import React, { Component, CSSProperties } from 'react';

type MenuButtonProps = {
  open: boolean,
  color?: string,
  onClick?: () => void,
};

type State = {
  open: boolean,
  color: string,
}


export default class MenuButton extends Component<MenuButtonProps, State> {
  constructor(props: MenuButtonProps) {
    super(props);
    this.state = {
      open: this.props.open ? this.props.open : false,
      color: this.props.color ? this.props.color : 'black',
    }
  }

  componentWillReceiveProps(nextProps: MenuButtonProps) {
    if (nextProps.open !== this.state.open) {
      this.setState({ open: nextProps.open });
    }
  }

  handleClick() {
    this.setState({ open: !this.state.open });
  }

  render() {
    return (
      <div style={style.container}
        onClick={this.props.onClick ? this.props.onClick :
          () => { this.handleClick(); }}>
        <div style={{
          ...style.line, background: this.state.color, ...style.lineTop,
          transform: this.state.open ? 'rotate(45deg)' : 'none',
        }} />
        <div style={{
          ...style.line, background: this.state.color, ...style.lineMiddle,
          opacity: this.state.open ? 0 : 1,
          transform: this.state.open ? 'translateX(-16px)' : 'none',
        }} />
        <div style={{
          ...style.line, background: this.state.color, ...style.lineBottom,
          transform: this.state.open ? 'translateX(-1px) rotate(-45deg)' : 'none',
        }} />
      </div>
    )
  }
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