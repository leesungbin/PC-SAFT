import { useUpdate } from 'react-three-fiber';
import React from 'react';
import {Vector3, Geometry} from 'three';

export function LineSegs({vertices, color}: {vertices: Vector3[], color?: string}) {
  const ref = useUpdate((geometry: Geometry) => {
    geometry.setFromPoints(vertices);
  }, []);
  const linecolor= color ? color : "black"
  return (
    <line onPointerOver={() => console.log('line')} onPointerOut={() => console.log('line out')}>
      <lineBasicMaterial attach="material" color={linecolor} linewidth={1}/>
      <geometry attach="geometry" ref={ref} />
    </line>
  )
}