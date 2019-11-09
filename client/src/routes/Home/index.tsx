import React from 'react';
import { Content } from '../../components/Content';
import { Canvas } from 'react-three-fiber';
import { Triangle } from '../../threeFragments/Triangle';
import { Vector3 } from 'three';
import { EQUIL_ENDPOINT } from '../../_lib/endpoint';

class Home extends React.Component {
  callEquil = async () => {
    const res = await fetch(EQUIL_ENDPOINT, {
      method: 'POST',
      mode: 'no-cors',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ T: 300, id: [20,40,50] }),
    });
    const json = await res.json()
    console.log(json);
  }
  
  render() {
    const trianglePoints = [new Vector3(0, 0, 0), new Vector3(1, 0, 0), new Vector3(1 / 2, Math.sqrt(3) / 2, 0)];
    return (
      <Content>
        <button onClick={() => this.callEquil()}>fetch test</button>
        <h1>Home</h1>
        <Canvas
          style={{ height: 500 }}
          camera={{ position: [1 / 2, Math.sqrt(3) / 4, 5], fov: 20 }}
        >
          <Triangle points={trianglePoints} />
        </Canvas>
      </Content>
    );
  }
}

export default Home;