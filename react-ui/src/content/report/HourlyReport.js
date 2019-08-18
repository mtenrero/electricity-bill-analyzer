import React from 'react';
import { BarChart, Bar, XAxis, YAxis, Legend, Tooltip } from 'recharts';

import './report.css';

class HourlyReport extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            report: props.report,
        }
    }

    render() {
        return(
            <div className="weekDaysReport">
                <BarChart data={this.state.report.hourlyReport.analysis}  width={500} height={300}>
                    <XAxis dataKey="hour"></XAxis>
                    <YAxis />
                    <Legend />
                    <Tooltip />
                    <Bar dataKey="mean" fill="#1890ff"></Bar>
                </BarChart>
            </div>
        );
    }
}

export default HourlyReport;