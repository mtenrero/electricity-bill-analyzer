import React from 'react';
import { BarChart, Bar, XAxis, YAxis, Legend, Tooltip } from 'recharts';

import { Tag } from 'antd';

import './report.css';

class WeekHourlyReport extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            report: props.report,
        }
    }

    render() {
        return(
            <div className="weekDaysReport">

                {this.state.report.weekdayHoursReport.map((value, index) => {
                    return (
                        <div>
                            <Tag color="#87d068">{value.WeekString}</Tag>
                            <BarChart data={value.report.analysis}  width={500} height={300}>
                                <XAxis dataKey="weekString"></XAxis>
                                <YAxis />
                                <Legend />
                                <Tooltip />
                                <Bar dataKey="mean" fill="#1890ff"></Bar>
                            </BarChart>
                        </div>
                    )
                })}

            </div>
        );
    }
}

export default WeekHourlyReport;