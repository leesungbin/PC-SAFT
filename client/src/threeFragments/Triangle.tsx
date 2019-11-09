import React from 'react';
import { Vector3 } from 'three';
import { LineSegs } from './LineSegs';

export function Triangle({ points }: { points: Vector3[]}) {
  const vertices = [...points, points[0]];
  return (
    <LineSegs vertices={vertices} />
  )
}