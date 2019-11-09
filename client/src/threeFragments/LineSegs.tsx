import { useUpdate } from 'react-three-fiber';
import React from 'react';
import {Vector3, Geometry} from 'three';

export function LineSegs({vertices}: {vertices: Vector3[]}) {
  const ref = useUpdate((geometry: Geometry) => {
    geometry.setFromPoints(vertices)
  }, []);

  return (
    <line>
      <geometry attach="geometry" ref={ref} />
      <lineBasicMaterial attach="material" color="black" />
    </line>
  )
}