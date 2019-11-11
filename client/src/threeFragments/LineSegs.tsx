import { useUpdate } from 'react-three-fiber';
import React from 'react';
import {Vector3, Geometry} from 'three';

export function LineSegs({vertices, color}: {vertices: Vector3[], color?: string}) {
  const ref = useUpdate((geometry: Geometry) => {
    geometry.setFromPoints(vertices)
  }, []);
  const linecolor= color ? color : "black"
  return (
    <line>
      <geometry attach="geometry" ref={ref} />
      <lineBasicMaterial attach="material" color={linecolor} />
    </line>
  )
}