import React from 'react';

type ContentProps = {
  children: any,
}
export function Content(props: ContentProps) {
  return (
    <div style={{marginLeft: '8%', marginRight: '8%'}}>
      {props.children}
    </div>
  )
}