import React from 'react';

type ContentProps = {
  children: any,
}
export function Content(props: ContentProps) {
  return (
    <div style={{marginLeft: 97, marginRight: 97}}>
      {props.children}
    </div>
  )
}