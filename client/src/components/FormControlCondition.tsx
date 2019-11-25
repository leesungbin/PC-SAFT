import React from 'react';
import { FormControlLabel, Checkbox, TextField } from '@material-ui/core';

type FormControlConditionProps = {
  valueDef: string,
  placeholder?: string,
  onChangeValue?: (e: number | undefined) => void,
  onError?: (err: string) => void,
};
type State = {
  const: boolean,
  val: string | null,
  hasErr?: boolean,
};
class FormControlCondition extends React.Component<FormControlConditionProps, State> {
  state: State = {
    const: false,
    val: null,
  }
  isNumber =(val: string) => {
    const floatRegex = new RegExp(/^(\d*\.?\d*)$/);
    return floatRegex.test(val);
  }
  render() {
    
    return (
      <div style={{ display: 'flex', alignItems: 'flex-end', justifyContent: 'space-between' }}>
        <FormControlLabel
          control={<Checkbox checked={this.state.const} onChange={() => { this.setState({ const: !this.state.const }) }} value={this.props.valueDef} />}
          label={this.props.valueDef}
        />
        <div style={{ marginBottom: 4 }}>
          <TextField
            error={this.state.hasErr}
            value={this.state.val? this.state.val : undefined}
            placeholder={this.props.placeholder}
            onChange={(e) => {
              const valStr = e.target.value;
              if (valStr.length === 0) {
                this.setState({val: null, hasErr: false});
                this.props.onChangeValue && this.props.onChangeValue(undefined)
                return
              }
              if (this.isNumber(valStr)) {
                const val = parseFloat(valStr);
                this.setState({ val: valStr, hasErr: false });
                this.props.onChangeValue && this.props.onChangeValue(val);
              } else {
                this.setState({hasErr: true});
                this.props.onError && this.props.onError('숫자를 입력해야합니다.');
              }
            }}
          />
        </div>
      </div>
    )
  }
}

export default FormControlCondition;