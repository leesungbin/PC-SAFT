import React from 'react';
import { Content } from '../../components/Content';
import { Canvas } from 'react-three-fiber';
import { Triangle } from '../../threeFragments/Triangle';
import { Vector3 } from 'three';
import { EQUIL_ENDPOINT, SEARCH_ENDPOINT } from '../../_lib/endpoint';
import Point from '../../threeFragments/Point';
import { TieLine } from '../../threeFragments/TieLine';

import ContinuosSlider from '../../components/ContinuosSlider';
import { Typography, ListItem, ListItemText, List, Chip } from '@material-ui/core';
import CalculatingIndicator from '../../components/CalculatingIndicator';
import SearchHeader from '../../components/SearchHeader';
import { ErrorSnack } from '../../components/Snack';

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
type Component = {
  id: string,
  name: string
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
  searchingName?: string,
  searchingLists?: any[],
  selectedComponents: Component[],
  error: string
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
  }

  callEquil = async () => {
    const { selectedComponents, error } = this.state;
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

  callSearch = async (searchingName: string) => {
    const res = await fetch(SEARCH_ENDPOINT, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ name: searchingName.toLowerCase() })
    });
    const json = await res.json();
    this.setState({ searchingLists: json.data });
    console.log(this.state.searchingLists);
  }
  selectComponent = (index: number) => {
    const { selectedComponents, searchingLists } = this.state;
    const comp: Component = { id: searchingLists![index].id, name: searchingLists![index].data.name };
    if (selectedComponents.length >= 3) {
      this.setState({ error: '3개까지만 계산할 수 있습니다.' });
      return;
    }
    this.setState({ selectedComponents: [...selectedComponents, comp], searchingName: '', searchingLists: [] });
  }
  cancleComponent = (index: number) => {
    const { selectedComponents } = this.state;
    const next = selectedComponents.filter((_, i) => { return i !== index });
    this.setState({ selectedComponents: next, error: '' });
  }
  render() {
    const trianglePoints = [new Vector3(0, 0, 0), new Vector3(1, 0, 0), new Vector3(1 / 2, Math.sqrt(3) / 2, 0)];

    const { data, waiting, T, P, mode, searchingLists } = this.state;
    const len = data.length;
    return (
      <div>
        <SearchHeader
          text={this.state.searchingName || ''}
          onChangeContent={(text) => {
            this.setState({ searchingName: text });
            text.length > 2 ? this.callSearch(text) : this.setState({ searchingLists: [] });
          }}
          listComponents={
            searchingLists && searchingLists.length > 0 ?
              <List style={{ maxHeight: 110, position: 'absolute', overflow: 'auto', top: 0, left: 0, width: '80%', paddingLeft: '10%', marginRight: '10%', opacity: 0.95, backgroundColor: 'white' }}>
                {searchingLists.map((list, i) => (
                  <ListItem button key={i} onClick={() => {
                    this.selectComponent(i);
                  }}>
                    <ListItemText primary={list.data.name} />
                  </ListItem>
                ))}
              </List> : <></>
          }
        />
        <div style={{ marginLeft: '10%', marginRight: '10%', height: 40, justifyContent: 'center' }}>
          {this.state.selectedComponents && this.state.selectedComponents.map((comp, i) => (
            <Chip style={{ marginRight: 10, marginBottom: 10 }} key={i} label={comp.name} variant="outlined" onDelete={() => this.cancleComponent(i)} />
          ))}
        </div>
        <Content>
          <div style={{ height: 100, display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
            <button onClick={() => this.callEquil()} style={styles.calculateButton}>Calculate</button>
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
        <div style={{ display: 'flex', justifyContent: 'center', zIndex: 0 }}>
          <Canvas
            style={{ height: this.props.height * 0.7, width: this.props.width }}
            camera={{ position: [1 / 2, Math.sqrt(3) / 4, 50], fov: 1.9, near: 1, far: -1 }}
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
        </div>
        <CalculatingIndicator open={waiting} />
        <ErrorSnack error={this.state.error} onClose={() => this.setState({ error: '' })} />
      </div>
    );
  }
}
const styles: { [key: string]: React.CSSProperties } = {
  calculateButton: {
    width: 100, height: 40, backgroundColor: "#FECF58", borderRadius: 10, fontSize: 17
  }
}

export default Home;