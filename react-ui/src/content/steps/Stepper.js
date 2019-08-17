import React from 'react';

import { Steps } from 'antd';

import './steps.css'
import Uploader from './Uploader';

const { Step } = Steps;

class Stepper extends React.Component {

    render() {
        return(
            <div className="Stepper">
                <Steps current={this.props.step}>
                    <Step title="Selecciona tu consumo" description="Localiza tu fichero de consumo" />
                    <Step title="Análisis" description="Espera mientras analizamos tus patrones de gasto energético" />
                    <Step title="Informe" description="Examina tus patrones de gasto energético" />
                </Steps>
                <Uploader></Uploader>
            </div>
        );
    }
    
}

export default Stepper;