import React from 'react';
import { Content } from '../../components/Content';
import './index.css'

class Document extends React.Component {
  render() {
    return (
      <Content>

        {/* <div className="box">
          Documents <br/>

          a
        </div> */}

        <div className="container">
        
          <h3 className = "title1">saft go 프로그램 작동 원리</h3>
          <p className = "para1">파란점 해석 </p> <br/><br/>
          
          <p className = "para2">분홍점 해석</p>

          
        </div>


      </Content>
    );
  }
}
export default Document;