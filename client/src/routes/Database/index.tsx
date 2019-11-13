import React from 'react';
import { Content } from '../../components/Content';
import { DATA_ENDPOINT } from '../../_lib/endpoint';
import DataTable from './Table';

type State = {
  columns: { title: string, field: string }[],
  data: object[]
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
    data: []
  }

  componentDidMount = async () => {
    const fetchData = await fetch(DATA_ENDPOINT, { method: 'POST' });
    const json = await fetchData.json();
    const data = json.data.map((e: any) => { return e.data });
  
    this.setState({ data });
  }

  render() {
    return (
      <Content>
        <div style={{marginTop: 20, maxWidth: '100%'}}>
          <DataTable
            columns={this.state.columns}
            data={this.state.data}
            title="DB Table"
          />
        </div>
      </Content>
    );
  }
}
export default Database;