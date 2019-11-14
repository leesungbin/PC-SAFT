import React, { useState } from 'react';
import { Slider } from '@material-ui/core';

type ContinuosSliderProps = {
  val: number,
  min?: number,
  max?: number,
  onChange?: (val: number) => void
}

export default function ContinuosSlider({ val, min, max, onChange }: ContinuosSliderProps) {
  const [value, setValue] = useState<number>(val);
  const handleChange = (event: any, newValue: number | number[]) => {
    setValue(newValue as number);
    console.log(newValue);
    onChange && onChange(newValue as number);
  }

  return <Slider step={0.001} value={value} onChange={handleChange} min={min} max={max} />
}