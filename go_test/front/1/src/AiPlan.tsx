import { Input, Button, Card, Tag, Tooltip } from 'antd';
// 🔥 全部替换为 Ant Design 官方合法图标
import {  BookOutlined, ClockCircleOutlined, ThunderboltOutlined } from '@ant-design/icons';
import { useState } from 'react';
import axios from 'axios';
import type { AxiosError } from 'axios';

interface PlanItem {
  name: string;
  description: string;
  priority: string;
}

interface AiPlanProps {
  onPlanGenerated?: (plans: PlanItem[]) => void;
}

export default function AiPlan({ onPlanGenerated }: AiPlanProps) {
  const [learningGoal, setLearningGoal] = useState("");
  const [learningDays, setLearningDays] = useState(7);
  const [loading, setLoading] = useState(false);
  const [planList, setPlanList] = useState<PlanItem[]>([]);

  const api = axios.create({
    baseURL: "http://localhost:8080",
  });

  const handleGenerateLearningPlan = async () => {
    if (!learningGoal) return;
    setLoading(true);

    try {
      const { data } = await api.post("/api/ai/generate-plan", {
        goal: learningGoal,
        days: learningDays,
      });
      setPlanList(data);
      if (onPlanGenerated) {
        onPlanGenerated(data);
      }
    } catch (err) {
      // 🔥 修复 TS 类型错误
      const error = err as AxiosError;
      alert("❌ 生成失败: " + (error.message || "未知错误"));
    } finally {
      setLoading(false);
    }
  };

  const getPriorityConfig = (priority: string) => {
    switch (priority) {
      case '高': return { color: 'red', bgColor: 'bg-red-50', borderColor: 'border-red-200', label: '高优先' };
      case '中': return { color: 'orange', bgColor: 'bg-orange-50', borderColor: 'border-orange-200', label: '中优先' };
      case '低': return { color: 'green', bgColor: 'bg-green-50', borderColor: 'border-green-200', label: '低优先' };
      default: return { color: 'gray', bgColor: 'bg-gray-50', borderColor: 'border-gray-200', label: priority };
    }
  };

  return (
    <div className="bg-gradient-to-br from-indigo-500 via-purple-500 to-pink-500 rounded-2xl p-6 shadow-2xl mb-6">
      {/* 标题区域 */}
      <div className="flex items-center gap-3 mb-6">
        <div className="w-12 h-12 bg-white/20 backdrop-blur-sm rounded-xl flex items-center justify-center">
          <BookOutlined className="text-white text-xl" />
        </div>
        <div>
          <h2 className="text-xl font-bold text-white">AI 学习计划生成器</h2>
          <p className="text-white/80 text-sm">智能生成个性化学习路径</p>
        </div>
      </div>

      {/* 输入区域 */}
      <div className="bg-white/10 backdrop-blur-sm rounded-xl p-4 mb-6">
        <div className="flex flex-wrap gap-3 items-center">
          <div className="flex-1 min-w-[200px]">
            <Input
              placeholder="输入学习目标，如：学会Go语言"
              value={learningGoal}
              onChange={(e) => setLearningGoal(e.target.value)}
              className="bg-white/90 border-0 focus:ring-2 focus:ring-white/50"
              prefix={<BookOutlined className="text-gray-400" />}
            />
          </div>
          <div className="flex items-center gap-2 bg-white/90 rounded-lg px-3 py-2">
            <ClockCircleOutlined className="text-gray-400" />
            <Input
              type="number"
              value={learningDays}
              onChange={(e) => setLearningDays(Math.max(1, Number(e.target.value)))}
              className="w-20 border-0 focus:ring-0 text-center font-semibold"
              min={1}
              max={30}
            />
            <span className="text-gray-500 text-sm">天</span>
          </div>
          <Button
            type="primary"
            onClick={handleGenerateLearningPlan}
            loading={loading}
            className="bg-white text-indigo-600 hover:bg-white/90 font-semibold shadow-lg hover:shadow-xl transition-all"
            icon={<ThunderboltOutlined className="text-indigo-600" />}
          >
            智能生成
          </Button>
        </div>
      </div>

      {/* 学习计划列表 */}
      {planList.length > 0 && (
        <div>
          <div className="flex items-center gap-2 mb-4">
            <BookOutlined className="text-white" />
            <h3 className="text-lg font-semibold text-white">📚 学习计划 ({planList.length}天)</h3>
          </div>
          <div className="space-y-3">
            {planList.map((item, index) => {
              const config = getPriorityConfig(item.priority);
              return (
                <Card 
                  key={index} 
                  className={`${config.bgColor} ${config.borderColor} border rounded-xl hover:shadow-lg transition-all duration-300 hover:-translate-y-1`}
                  bordered={false}
                >
                  <div className="flex items-start gap-4">
                    {/* 天数标识 */}
                    <div className="flex-shrink-0 w-10 h-10 bg-gradient-to-br from-indigo-500 to-purple-500 rounded-xl flex items-center justify-center text-white font-bold shadow-md">
                      {index + 1}
                    </div>
                    {/* 内容 */}
                    <div className="flex-1 min-w-0">
                      <h4 className="font-semibold text-gray-800 mb-1 flex items-center gap-2">
                        {item.name}
                        <Tooltip title={`${config.label}任务`}>
                          <Tag color={config.color} className="text-xs">{config.label}</Tag>
                        </Tooltip>
                      </h4>
                      <p className="text-gray-600 text-sm leading-relaxed">{item.description}</p>
                    </div>
                  </div>
                </Card>
              );
            })}
          </div>
        </div>
      )}
    </div>
  );
}