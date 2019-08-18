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
                <Card title="Media de consumo por días de la semana" className="centerReport">
                    <WeekDaysReport {...this.props}></WeekDaysReport> 
                </Card>
                
            </div>
        )
    }
}

export default Report;