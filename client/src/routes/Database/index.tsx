import React from 'react';
import { DATA_ENDPOINT } from '../../_lib/endpoint';
import DataTable from './Table';
import './index.css';

type Component = {
  id: string,
  name: string,
  d: number,
  e: number,
}
type State = {
  columns: { title: string, field: string }[],
  data: Component[],
  isDZeroFilterOn: boolean,
  isDNonZeroFilterOn: boolean,
  isEZeroFilterOn: boolean,
  isENonZeroFilterOn: boolean,
}
class Database extends React.Component<{}, State> {
  state: State = {
    columns: [
      { title: 'Name', field: 'name' },
      { title: 'Mw', field: 'Mw' },
      { title: 'Pc', field: 'Pc' },
      { title: 'Tc', field: 'Tc' },
      { title: 'Tb', field: 'Tb' },
      { title: 'ω', field: 'w' },
      { title: 'ε', field: 'epsilon' },
      { title: 'm', field: 'm' },
      { title: 'σ', field: 'sigma' },
      { title: 'k', field: 'k' },
      { title: 'e', field: 'e' },
      { title: 'd', field: 'd' },
      { title: 'x', field: 'x' },
    ],
    data: [],
    isDZeroFilterOn: false,
    isDNonZeroFilterOn: false,
    isEZeroFilterOn: false,
    isENonZeroFilterOn: false,
  }

  componentDidMount = async () => {
    const fetchData = await fetch(DATA_ENDPOINT, { method: 'POST' });
    const json = await fetchData.json();
    const data = json.data.map((e: any) => { return e.data });

    this.setState({ data });
  }

  filterData(data: Component[]) {
    if (this.state.isDZeroFilterOn) {
      return data.filter((datum: Component) => datum.d === 0)
    } else if (this.state.isDNonZeroFilterOn) {
      return data.filter((datum: Component) => datum.d !== 0)
    } else if (this.state.isEZeroFilterOn) {
      return data.filter((datum: Component) => datum.e === 0)
    } else if (this.state.isENonZeroFilterOn) {
      return data.filter((datum: Component) => datum.e !== 0)
    }
    return data
  }

  render() {
    return (
      <div style={{ margin: 20, maxWidth: '100%' }}>
        <button className={'btn btn-1'} style={{ background: this.state.isDZeroFilterOn ? 'skyblue' : 'white' }} onClick={() => this.setState({ isDZeroFilterOn: !this.state.isDZeroFilterOn })}>
          <span><strong>D가 0인 것만 보기 필터 켜기</strong></span>
        </button>
        <button className={'btn btn-1'} style={{ background: this.state.isDNonZeroFilterOn ? 'skyblue' : 'white' }} onClick={() => this.setState({ isDNonZeroFilterOn: !this.state.isDNonZeroFilterOn })}>
          <span><strong>D가 0이 아닌 것만 보기 필터 켜기</strong></span>
        </button>
        <button className={'btn btn-1'} style={{ background: this.state.isEZeroFilterOn ? 'skyblue' : 'white' }} onClick={() => this.setState({ isEZeroFilterOn: !this.state.isEZeroFilterOn })}>
          <span><strong>E가 0인 것만 보기 필터 켜기</strong></span>
        </button>
        <button className={'btn btn-1'} style={{ background: this.state.isENonZeroFilterOn ? 'skyblue' : 'white' }} onClick={() => this.setState({ isENonZeroFilterOn: !this.state.isENonZeroFilterOn })}>
          <span><strong>E가 0이 아닌 것만 보기 필터 켜기</strong></span>
        </button>
        <DataTable
          columns={this.state.columns}
          data={this.filterData(this.state.data)}
          title="DB Table"
        />
      </div>
    );
  }
}
export default Database;