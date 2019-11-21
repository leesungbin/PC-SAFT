import React from 'react';
import { Content } from '../../components/Content';
import { Canvas } from 'react-three-fiber';
import { Triangle } from '../../threeFragments/Triangle';
import { Vector3 } from 'three';
import { EQUIL_ENDPOINT, DATA_ENDPOINT } from '../../_lib/endpoint';
import Point from '../../threeFragments/Point';
import { TieLine } from '../../threeFragments/TieLine';

import ContinuosSlider from '../../components/ContinuosSlider';
import { Typography, ListItem, ListItemText, List, Chip } from '@material-ui/core';
import CalculatingIndicator from '../../components/CalculatingIndicator';
import SearchHeader from '../../components/SearchHeader';
import { ErrorSnack } from '../../components/Snack';

import './index.css';
import ComponentSelector from '../../components/ComponentSelector';

type FetchResult = {
  result: {
    data: { P: number, T: number, x: number[], y: number[] }[],
    header: { min: number, max: number },
    mode: string,
    names: string[],
  }
}
type HomeProps = {
  width: number,
  height: number,
}
export type Component = {
  id: string,
  name: string,
  selected?: boolean,
}
type State = {
  data: {
    P: number,
    T: number,
    x: number[],
    y: number[],
  }[],
  waiting: boolean,
  T: number,
  P: number,
  openSelector: boolean,
  mode?: string,
  min?: number,
  max?: number,
  names?: string[],
  selectedComponents: Component[],
  error: string,
  components: Component[],
}

class Home extends React.Component<HomeProps, State> {
  state: State = {
    data: [],
    waiting: false,
    T: 300,
    P: 1,
    openSelector: false,
    selectedComponents: [],
    error: '',
    components: [],
  }

  componentDidMount = async () => {
    const fetchData = await fetch(DATA_ENDPOINT, { method: 'POST' });
    const json = await fetchData.json();
    const components = json.data.map((e: any) => { return {...e.data, id: e.id} });

    this.setState({ components });
  }
  componentDidUpdate = (_: HomeProps, prevState: State) => {
    if (prevState.components !== this.state.components) {
      const selectedComponents = this.state.components.filter(component => component.selected === true);
      this.setState({ selectedComponents });
    }
  }
  callEquil = async () => {
    const { selectedComponents, error } = this.state;
    console.log(selectedComponents);
    if (selectedComponents.length === 3 && error === '') {
      const id = selectedComponents.map(comp => { return comp.id });

      this.setState({ waiting: true });
      const res = await fetch(EQUIL_ENDPOINT, {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ P: 1.013, id }),
      });
      const json: FetchResult = await res.json()
      const { data, header, mode, names } = json.result;
      if (data.length === 0) {
        this.setState({ error: '계산이 잘 되지 않는 조합입니다...' });
        return;
      }
      const { min, max } = header;
      this.setState({ data, min, max, mode, names, waiting: false });
      return;
    } else {
      this.setState({ error: '아직 계산을 시작 할 수 없습니다.' })
      return;
    }
  }
  cancleComponent = (index: number) => {
    let components = [...this.state.components]
    const cancleComponentIdx = this.state.components.findIndex(component => component.name === this.state.selectedComponents[index].name);
    if (cancleComponentIdx > -1) {
      components[cancleComponentIdx].selected = false;
      this.setState({ components });
    }
  }

  render() {
    const trianglePoints = [new Vector3(0, 0, 0), new Vector3(1, 0, 0), new Vector3(1 / 2, Math.sqrt(3) / 2, 0)];

    const { data, waiting, T, P, mode } = this.state;
    const len = data.length;
    return (
      <div>
        <div style={{ marginTop: 10, marginLeft: '10%', marginRight: '10%', height: 40, justifyContent: 'center' }}>
          {this.state.selectedComponents && this.state.selectedComponents.map((comp, i) => (
            <Chip style={{ marginRight: 10, marginBottom: 10 }} key={i} label={comp.name} variant="outlined" onDelete={() => this.cancleComponent(i)} />
          ))}
        </div>
        <div style={{ display: 'flex', justifyContent: 'center', zIndex: 0, flexWrap: "wrap" }}>
          <div style={{ width: 300 }}>
            <ComponentSelector
              components={this.state.components}
              onClickContent={(i) => {
                let newComponents = [...this.state.components];
                // switching value
                newComponents[i].selected = newComponents[i].selected ? false : true;
                this.setState({ components: newComponents });
              }}
              calcButton = {<button onClick={() => this.callEquil()} style={styles.calculateButton}>Calculate</button>}
            />
          </div>
          <div>
            <Canvas
              style={{ height: this.props.height * 0.7, width: this.props.width * 0.7 }}
              camera={{ position: [1 / 2, Math.sqrt(3) / 4, 3], fov: 18, near: 1, far: -1 }}
            >
              <mesh rotation={[0, 0, 0]}>
                <Triangle points={trianglePoints} />
                {len && data.map((e, i) => {
                  if (mode === "BUBLP" && e.P < P * 1.01 && e.P > P * 0.99) {
                    return (
                      <mesh key={i}>
                        <Point abc={e.x} val={0} t={0} />
                        <TieLine x={e.x} y={e.y} val={0} color="green" />
                        <Point abc={e.y} val={0} t={1} />
                      </mesh>
                    )
                  }
                  else if (mode === "BUBLT" && e.T < T * 1.005 && e.T > T * 0.995) {
                    return (
                      <mesh key={i}>
                        <Point abc={e.x} val={0} t={0} />
                        <TieLine x={e.x} y={e.y} val={0.1} color="green" />
                        <Point abc={e.y} val={0} t={1} />
                      </mesh>
                    )
                  }
                  return <></>
                })}
              </mesh>
            </Canvas>
            <Content>
              <div style={{ height: 100, display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
                <div style={{ display: 'flex', width: '100%', flexDirection: 'row', justifyContent: 'space-around' }}>
                  {mode && mode === "BUBLP" ?
                    <div style={{ flex: 1, padding: '0,10%,10%,0' }}>
                      <Typography gutterBottom>P : {P.toFixed(3)} atm</Typography>
                      <ContinuosSlider val={P} onChange={(P) => this.setState({ P })} min={this.state.min} max={this.state.max} />
                    </div>
                    : mode === "BUBLT" ?
                      <div style={{ flex: 1, padding: '0,10%,10%,0' }}>
                        <Typography gutterBottom>T : {T.toFixed(3)} K</Typography>
                        <ContinuosSlider val={T} onChange={(T) => this.setState({ T })} min={this.state.min} max={this.state.max} />
                      </div>
                      : <></>
                  }
                </div>
              </div>
            </Content>
          </div>
        </div>
        <CalculatingIndicator open={waiting} />
        <ErrorSnack error={this.state.error} onClose={() => this.setState({ error: '' })} />
      </div>
    );
  }
}
const styles: { [key: string]: React.CSSProperties } = {
  calculateButton: {
    width: '100%', height: 40, backgroundColor: "#FECF58", borderRadius: 10, fontSize: 17
  }
}

export default Home;