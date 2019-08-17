import React from 'react';

import { Layout } from 'antd';
import './header.css';


const { Header } = Layout;

function HeaderLayout(props) {
    return(
        <Header>
            <div className="logo">electricity bill analyzer</div>
            <div className="logo_extra">by mtenrero.com</div>
    
        </Header>
    );
}

export default HeaderLayout;