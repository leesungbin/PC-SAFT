import React from 'react';
import { Content } from '../../components/Content';
import { Canvas } from 'react-three-fiber';
import { Triangle } from '../../threeFragments/Triangle';
import { Vector3 } from 'three';
import { EQUIL_ENDPOINT } from '../../_lib/endpoint';
import Point from '../../threeFragments/Point';

type State = {
  data: {
    P: number,
    T: number,
    x: number[],
    y: number[],
  }[]
}
class Home extends React.Component<{}, State> {
  state: State = {
    data: []
  }

  callEquil = async () => {
    const res = await fetch(EQUIL_ENDPOINT, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ T: 300, id: [20, 40, 50] }),
    });
    const json = await res.json()
    this.setState({ data: json.data });
  }

  render() {
    const trianglePoints = [new Vector3(0, 0, 0), new Vector3(1, 0, 0), new Vector3(1 / 2, Math.sqrt(3) / 2, 0)];
    const { data } = this.state;

    return (
      <Content>
        <button onClick={() => this.callEquil()}>fetch test</button>
        <h1>Home</h1>
        <Canvas
          style={{ height: 500 }}
          camera={{ position: [1 / 2, Math.sqrt(3) / 4, 500], fov: 0.2 }}
        >
          <mesh rotation={[0, 0, 0]}>
            <Triangle points={trianglePoints} />
            {data && data.map(e => (
              <>
                <Point abc={[...e.x]} val={e.P} t={0} />
                <Point abc={[...e.y]} val={e.P} t={1} />
              </>
            ))}
          </mesh>
        </Canvas>
      </Content>
    );
  }
}

export default Home;