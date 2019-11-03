import React, { useRef } from 'react';
import { Canvas, useFrame } from 'react-three-fiber';

const App: React.FC = () => {
  return (
    <div className="App">
      <Canvas>
        <Thing />
      </Canvas>
    </div>
  );
}


export function Thing() {
  const ref = useRef<any>()
  useFrame(() => [(ref.current.rotation.x +=0.08), (ref.current.rotation.y += 0.05)])
  return (
    <mesh
      ref={ref}
      onClick={e => console.log('click')}
      onPointerOver={e => console.log('hover')}
      onPointerOut={e => console.log('unhover')}>
      <boxBufferGeometry attach="geometry" args={[3, 3, 0.2]} />
      <meshNormalMaterial attach="material" />
      <lineBasicMaterial/>
    </mesh>
  )
}
export default App;
