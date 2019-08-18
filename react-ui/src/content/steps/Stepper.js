import React from 'react';

import { Steps, message } from 'antd';

import './steps.css'
import Uploader from './Uploader';

const { Step } = Steps;

class Stepper extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            step: props.step,
            servererror: "process",
        }
    }

    moveStep() {
        this.setState({step: this.state.step + 1})
    }

    handleError() {
        message.error("Error procesando el fichero, posiblemente no tenga un contenido válido")
        this.setState({servererror: "error"})
    }

    render() {
        return(
            <div className="Stepper">
                <Steps current={this.state.step} status={this.state.servererror}>
                    <Step title="Selecciona tu consumo" description="Localiza tu fichero de consumo" />
                    <Step title="Análisis" description="Espera mientras analizamos tus patrones de gasto energético" />
                    <Step title="Informe" description="Examina tus patrones de gasto energético" />
                </Steps>

                { this.state.step < 2 ? <Uploader next={this.moveStep.bind(this)} servererror={this.handleError.bind(this)} step={this.state.step} onSubmit={this.handleSubmit}></Uploader>: null}
            </div>
        );
    }
    
}

export default Stepper;