import React from 'react';
import WeekDaysReport from './WeekDaysReport';
import { Card } from 'antd';

class Report extends React.Component {
    
    constructor(props) {
        super(props);
        this.state = {
            report: props.report,
        }
    }

    render() {
        return(
            <div className="report">
                <Card title="Consumo por días de la semana">
                    <WeekDaysReport {...this.props}></WeekDaysReport> 
                </Card>
                
            </div>
        )
    }
}

export default Report;