import React from 'react';

import { Steps, Form, message } from 'antd';

import './steps.css'
import Uploader from './Uploader';

const { Step } = Steps;

class Stepper extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            step: props.step
        }
    }

    moveStep(step) {
        message.info("Step")
        this.setState({step: this.state.step + 1})
    }

    render() {
        const FormUpload = Form.create({}) (Uploader);
        return(
            <div className="Stepper">
                <Steps current={this.state.step}>
                    <Step title="Selecciona tu consumo" description="Localiza tu fichero de consumo" />
                    <Step title="Análisis" description="Espera mientras analizamos tus patrones de gasto energético" />
                    <Step title="Informe" description="Examina tus patrones de gasto energético" />
                </Steps>
                <FormUpload next={this.moveStep.bind(this)}></FormUpload>
            </div>
        );
    }
    
}

export default Stepper;