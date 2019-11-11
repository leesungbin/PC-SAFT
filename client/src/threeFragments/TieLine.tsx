import React from 'react';
import { coordChange } from '../_lib/coordChange';
import { LineSegs } from './LineSegs';
import { Vector3 } from 'three';

export function TieLine({x,y,val,color}: {x: number[], y: number[], val: number, color?: string}) {
  const liq = coordChange(x[0], x[1], x[2]);
  const vap = coordChange(y[0], y[1], y[2]);
  const X = new Vector3(liq.x, liq.y, val);
  const Y = new Vector3(vap.x, vap.y, val);
  return (
    <LineSegs vertices={[X,Y]} color={color}/>
  )
}