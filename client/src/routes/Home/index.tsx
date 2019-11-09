import React from 'react';
import { Content } from '../../components/Content';
import { Canvas } from 'react-three-fiber';
import { Triangle } from '../../threeFragments/Triangle';
import { Vector3 } from 'three';

class Home extends React.Component {

  render() {
    const trianglePoints=[new Vector3(0,0,0), new Vector3(1,0,0), new Vector3(1/2, Math.sqrt(3)/2, 0)];

    return (
      <Content>
        <h1>Home</h1>
        <Canvas
          style={{ height: 500 }}
          camera={{ position: [1/2, Math.sqrt(3)/4, 5], fov: 20 }}
          >
            <Triangle points={trianglePoints}/>
        </Canvas>
      </Content>
    );
  }
}

export default Home;