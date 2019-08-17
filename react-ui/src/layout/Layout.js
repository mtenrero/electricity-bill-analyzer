import React from 'react';

import { Layout } from 'antd';

import HeaderLayout from './Header'
import Foot from'./Footer'

function LayoutElec(props) {
    return(
        <div className="App">
            <Layout className="layout">
                <HeaderLayout></HeaderLayout>
                <Foot></Foot>
            </Layout>
            
        </div>
    );
}

export default LayoutElec;