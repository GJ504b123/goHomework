import { useEffect, useState } from "react";
import axios from "axios";
import { Input, Button, Select, Tag, Checkbox, Card, Space } from "antd";
import { CheckCircleOutlined, CiCircleOutlined, ToolOutlined } from '@ant-design/icons';
import AiPlan from "./AiPlan";

interface Task {
  id: number;
  name: string;
  description: string;
  status: string;
  priority: string;
  created_at: string;
}

interface PlanItem {
  name: string;
  description: string;
  priority: string;
}

const api = axios.create({
  baseURL: "http://localhost:8080/api",
});

export default function TaskList() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [status, setStatus] = useState("待办");
  const [priority, setPriority] = useState("中");
  const [filterStatus, setFilterStatus] = useState("全部");
  const [filterPriority, setFilterPriority] = useState("全部");

  const handleAddTask = async () => {
    if (!name) return;
    await api.post("/tasks", { name, description, status, priority });
    setName("");
    setDescription("");
    setStatus("待办");
    setPriority("中");
    fetchTask();
  };

  const handleAddPlanToTasks = async (plans: PlanItem[]) => {
    for (const plan of plans) {
      await api.post("/tasks", {
        name: plan.name,
        description: plan.description,
        status: "待办",
        priority: plan.priority,
      });
    }
    fetchTask();
    alert(`✅ 已成功添加 ${plans.length} 个学习任务！`);
  };

  const fetchTask = async () => {
    const res = await api.get("/tasks");
    setTasks(res.data);
  };

  const fetchTaskByStatus = async (status: string) => {
    const res = status === "全部" 
      ? await api.get("/tasks") 
      : await api.get(`/tasks/status/${status}`);
    setTasks(res.data);
  };

  const fetchTaskByPriority = async (priority: string) => {
    const res = priority === "全部" 
      ? await api.get("/tasks") 
      : await api.get(`/tasks/priority/${priority}`);
    setTasks(res.data);
  };

  const handleDelete = async (id: number) => {
    await api.delete(`/tasks/${id}`);
    fetchTask();
  };

  const updateTask = async (task: Task) => {
    const newStatus = task.status === "已完成" ? "待办" : "已完成";
    await api.put(`/tasks/${task.id}`, { ...task, status: newStatus });
    fetchTask();
  };

  const getStatusConfig = (status: string) => {
    return status === "已完成" 
      ? { color: 'green', label: '已完成', icon: <CheckCircleOutlined className="text-green-500" /> }
      : { color: 'orange', label: '待办', icon: <CiCircleOutlined className="text-orange-400" /> };
  };

  const getPriorityConfig = (priority: string) => {
    switch (priority) {
      case '高': return { color: 'red', bgColor: 'bg-red-50', text: '高优先级' };
      case '中': return { color: 'orange', bgColor: 'bg-orange-50', text: '中优先级' };
      case '低': return { color: 'green', bgColor: 'bg-green-50', text: '低优先级' };
      default: return { color: 'gray', bgColor: 'bg-gray-50', text: priority };
    }
  };

  const formatDate = (dateStr: string) => {
    const date = new Date(dateStr);
    return date.toLocaleDateString('zh-CN', { 
      month: 'short', 
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  };

  useEffect(() => {
    fetchTask();
  }, []);

  const pendingCount = tasks.filter(t => t.status !== "已完成").length;

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 via-blue-50 to-indigo-50 p-4 md:p-8">
      <div className="max-w-4xl mx-auto">
        {/* 页面标题 */}
        <div className="text-center mb-8">
          <div className="inline-flex items-center gap-3 mb-2">
            <div className="w-12 h-12 bg-gradient-to-br from-indigo-500 to-purple-500 rounded-2xl flex items-center justify-center shadow-lg">
              <ToolOutlined className="text-white text-xl" />
            </div>
            <h1 className="text-3xl font-bold bg-gradient-to-r from-indigo-600 to-purple-600 bg-clip-text text-transparent">
              任务管理系统
            </h1>
          </div>
          <p className="text-gray-500">高效管理你的学习与工作任务</p>
          {pendingCount > 0 && (
            <div className="mt-3 inline-flex items-center gap-2 bg-orange-100 text-orange-700 px-4 py-1.5 rounded-full text-sm font-medium">
              <CiCircleOutlined className="text-orange-500" />
              {pendingCount} 个待办任务
            </div>
          )}
        </div>

        {/* AI 学习计划生成器 */}
        <AiPlan onPlanGenerated={handleAddPlanToTasks} />

        {/* 添加任务区域 */}
        <Card className="mb-6 border-0 shadow-lg rounded-2xl bg-white">
          <div className="flex items-center gap-2 mb-4">
            <h3 className="font-semibold text-gray-800">添加新任务</h3>
          </div>
          <div className="flex flex-wrap gap-3">
            <Input
              placeholder='任务名称'
              value={name}
              onChange={(e) => setName(e.target.value)}
              className="flex-1 min-w-[150px]"
              size="large"
            />
            <Input
              placeholder='任务描述（可选）'
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              className="flex-1 min-w-[150px]"
              size="large"
            />
            <Select 
              defaultValue="待办" 
              options={[
                { value: '待办', label: "待办" },
                { value: '已完成', label: "已完成" },
              ]}
              value={status}
              onChange={(value) => setStatus(value)}
              size="large"
              className="w-28"
            />
            <Select 
              defaultValue="中" 
              options={[
                { value: '高', label: "高" },
                { value: '中', label: "中" },
                { value: '低', label: "低" },
              ]}
              value={priority}
              onChange={(value) => setPriority(value)}
              size="large"
              className="w-28"
            />
            <Button 
              type="primary" 
              onClick={handleAddTask}
              size="large"
              className="bg-gradient-to-r from-indigo-500 to-purple-500 hover:from-indigo-600 hover:to-purple-600 shadow-md hover:shadow-lg transition-all"
              disabled={!name}
            >
              添加任务
            </Button>
          </div>
        </Card>

        {/* 筛选区域 */}
        <Card className="mb-6 border-0 shadow-lg rounded-2xl bg-white">
          <div className="flex items-center gap-3">
            <span className="text-gray-600 font-medium">筛选：</span>
            <Select 
              defaultValue="全部" 
              options={[
                { value: '全部', label: "全部状态" },
                { value: '待办', label: "待办" },
                { value: '已完成', label: "已完成" },
              ]}
              value={filterStatus}
              onChange={(value) => { setFilterStatus(value); fetchTaskByStatus(value); }}
              className="w-32"
            />
            <Select 
              defaultValue="全部" 
              options={[
                { value: '全部', label: "全部优先级" },
                { value: '高', label: "高优先级" },
                { value: '中', label: "中优先级" },
                { value: '低', label: "低优先级" },
              ]}
              value={filterPriority}
              onChange={(value) => { setFilterPriority(value); fetchTaskByPriority(value); }}
              className="w-32"
            />
            <div className="flex-1" />
            <span className="text-gray-400 text-sm">共 {tasks.length} 个任务</span>
          </div>
        </Card>

        {/* 任务列表 */}
        <div className="space-y-3">
          {tasks.length === 0 ? (
            <Card className="border-0 shadow-lg rounded-2xl bg-white text-center py-12">
              <div className="w-20 h-20 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4">
                <ToolOutlined className="text-gray-300 text-3xl" />
              </div>
              <h3 className="text-lg font-medium text-gray-600 mb-2">暂无任务</h3>
              <p className="text-gray-400">添加一个新任务开始你的计划吧！</p>
            </Card>
          ) : (
            tasks.map((task) => {
              const statusConfig = getStatusConfig(task.status);
              const priorityConfig = getPriorityConfig(task.priority);
              const isCompleted = task.status === "已完成";
              
              return (
                <Card 
                  key={task.id} 
                  className={`border-0 shadow-md hover:shadow-lg transition-all duration-300 rounded-2xl ${isCompleted ? 'bg-gray-50' : 'bg-white'}`}
                >
                  <div className="flex items-start gap-4">
                    {/* 复选框 */}
                    <Checkbox 
                      checked={isCompleted}
                      onChange={() => updateTask(task)}
                      className="mt-1"
                    >
                      {statusConfig.icon}
                    </Checkbox>
                    
                    {/* 任务内容 */}
                    <div className="flex-1 min-w-0">
                      <h4 className={`font-semibold ${isCompleted ? 'text-gray-400 line-through' : 'text-gray-800'} mb-1`}>
                        {task.name}
                      </h4>
                      {task.description && (
                        <p className={`text-sm ${isCompleted ? 'text-gray-300' : 'text-gray-500'} mb-2`}>
                          {task.description}
                        </p>
                      )}
                      <Space size="small">
                        <Tag color={statusConfig.color} className="text-xs">
                          {statusConfig.label}
                        </Tag>
                        <Tag className={`text-xs ${priorityConfig.bgColor} text-gray-600 border-0`}>
                          {priorityConfig.text}
                        </Tag>
                        <span className="text-xs text-gray-400">
                          {formatDate(task.created_at)}
                        </span>
                      </Space>
                    </div>
                    
                    {/* 删除按钮 */}
                    <Button 
                      type="text" 
                      danger 
                      onClick={() => handleDelete(task.id)}
                      className="opacity-0 group-hover:opacity-100 transition-opacity"
                    />
                  </div>
                </Card>
              );
            })
          )}
        </div>
      </div>
    </div>
  );
}
