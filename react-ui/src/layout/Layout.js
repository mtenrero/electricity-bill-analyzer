import React from 'react';

import { Layout } from 'antd';

import HeaderLayout from './Header'
import Foot from './Footer'
import Info from '../content/info/Info'
import Stepper from '../content/steps/Stepper';

const { Content } = Layout;

function LayoutElec(props) {
    return(
        <div className="App">
            <Layout className="layout">
                <HeaderLayout></HeaderLayout>
                <Content style={{ padding: '0 50px' }}>
                    <div style={{ background: '#fff', padding: 24, minHeight: 280 }}>
                        <Info></Info>
                        <br></br>
                        <Stepper step={0}></Stepper>
                    </div>
                </Content>
                <Foot></Foot>
            </Layout>
            
        </div>
    );
}

export default LayoutElec;