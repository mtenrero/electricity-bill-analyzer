import React from 'react';
import { BarChart, Bar, XAxis, YAxis, Legend, Tooltip } from 'recharts';

import './report.css';

class WeekDaysReport extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            report: props.report,
        }
    }

    render() {
        return(
            <div className="weekDaysReport">
                <BarChart data={this.state.report.weekDaysReport.analysis}  width={500} height={300}>
                    <XAxis dataKey="weekDay"></XAxis>
                    <YAxis />
                    <Legend />
                    <Tooltip />
                    <Bar dataKey="mean" fill="#82ca9d"></Bar>
                </BarChart>
            </div>
        );
    }
}

export default WeekDaysReport;