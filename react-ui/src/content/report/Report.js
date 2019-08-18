import React from 'react';
import WeekDaysReport from './WeekDaysReport';
import { Card } from 'antd';
import HourlyReport from './HourlyReport';

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
                <Card title="Media de consumo por días de la semana" className="cardReport">
                    <WeekDaysReport {...this.props}></WeekDaysReport> 
                </Card>
                <Card title="Media de consumo horario" className="cardReport">
                    <HourlyReport {...this.props}></HourlyReport>
                </Card>
                
            </div>
        )
    }
}

export default Report;