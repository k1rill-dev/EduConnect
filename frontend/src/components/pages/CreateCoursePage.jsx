import React, { useState } from "react";
import { Button, Label, TextInput, FileInput, Accordion } from "flowbite-react";

const CreateCoursePage = () => {
  const [courseData, setCourseData] = useState({
    title: "",
    startDate: "",
    endDate: "",
    topics: [],
  });

  const [newTopic, setNewTopic] = useState("");
  const [assignments, setAssignments] = useState({});
  const [newAssignments, setNewAssignments] = useState({});
  const [editingAssignmentIndex, setEditingAssignmentIndex] = useState(null);
  const [editingAssignmentData, setEditingAssignmentData] = useState({ title: "", theoryFile: null });

  const handleCourseChange = (e) => {
    const { name, value } = e.target;
    setCourseData((prev) => ({ ...prev, [name]: value }));
  };

  const handleAddAssignment = (topicIndex) => {
    const assignment = newAssignments[topicIndex];
    if (assignment && assignment.title.trim()) {
      setAssignments((prev) => ({
        ...prev,
        [topicIndex]: [
          ...(prev[topicIndex] || []),
          { ...assignment },
        ],
      }));

      setNewAssignments((prev) => ({
        ...prev,
        [topicIndex]: { title: "", theoryFile: null },
      }));
    }
  };

  const handleAddTopic = () => {
    if (newTopic.trim()) {
      setCourseData((prev) => ({
        ...prev,
        topics: [...prev.topics, { title: newTopic }],
      }));
      setNewTopic("");

      setNewAssignments((prev) => ({
        ...prev,
        [courseData.topics.length]: { title: "", theoryFile: null },
      }));
    }
  };

  const handleEditAssignment = (topicIndex, assignmentIndex) => {
    setEditingAssignmentIndex({ topicIndex, assignmentIndex });
    const assignment = assignments[topicIndex][assignmentIndex];
    setEditingAssignmentData({ ...assignment });
  };

  const handleSaveAssignmentEdit = () => {
    const { topicIndex, assignmentIndex } = editingAssignmentIndex;
    setAssignments((prev) => {
      const updatedAssignments = { ...prev };
      updatedAssignments[topicIndex][assignmentIndex] = { ...editingAssignmentData };
      return updatedAssignments;
    });
    setEditingAssignmentIndex(null);
    setEditingAssignmentData({ title: "", theoryFile: null });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    const courseWithAssignments = {
      ...courseData,
      topics: courseData.topics.map((topic, index) => ({
        ...topic,
        assignments: assignments[index] || [],
      })),
    };
    console.log("Созданный курс:", courseWithAssignments);
    alert("Курс успешно создан!");
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="bg-gradient-to-r from-green-500 to-blue-600 text-white py-10 shadow-md">
        <div className="container mx-auto px-6 text-center">
          <h1 className="text-4xl font-bold">Создать курс</h1>
          <p className="mt-2 text-lg">Заполните форму ниже, чтобы опубликовать новый курс.</p>
        </div>
      </div>

      <div className="container mx-auto px-6 py-10">
        <form onSubmit={handleSubmit} className="bg-white p-6 rounded-lg shadow-lg space-y-6">
          <div>
            <Label htmlFor="title" className="mb-2 block">Название курса</Label>
            <TextInput
              id="title"
              name="title"
              type="text"
              placeholder="Введите название курса"
              value={courseData.title}
              onChange={handleCourseChange}
              required
            />
          </div>
          <div>
            <Label htmlFor="startDate" className="mb-2 block">Дата начала курса</Label>
            <TextInput
              id="startDate"
              name="startDate"
              type="date"
              value={courseData.startDate}
              onChange={handleCourseChange}
              required
            />
          </div>
          <div>
            <Label htmlFor="endDate" className="mb-2 block">Дата окончания курса</Label>
            <TextInput
              id="endDate"
              name="endDate"
              type="date"
              value={courseData.endDate}
              onChange={handleCourseChange}
              required
            />
          </div>

          <div>
            <h2 className="text-lg font-semibold mb-4">Темы и задания</h2>

            {courseData.topics.map((topic, topicIndex) => (
              <Accordion key={topicIndex} className="mb-4">
                <Accordion.Panel>
                  <Accordion.Title>{topic.title}</Accordion.Title>
                  <Accordion.Content>
                    {(assignments[topicIndex] || []).map((assignment, idx) => (
                      <div key={idx} className="mb-4 p-4 border rounded shadow-sm">
                        {editingAssignmentIndex &&
                        editingAssignmentIndex.topicIndex === topicIndex &&
                        editingAssignmentIndex.assignmentIndex === idx ? (
                          <>
                            <Label htmlFor="editTitle" className="mb-2 block">Название задания</Label>
                            <TextInput
                              id="editTitle"
                              type="text"
                              value={editingAssignmentData.title}
                              onChange={(e) =>
                                setEditingAssignmentData((prev) => ({ ...prev, title: e.target.value }))
                              }
                            />
                            <Label htmlFor="editFile" className="mb-2 block">Файл теории</Label>
                            <FileInput
                              id="editFile"
                              onChange={(e) =>
                                setEditingAssignmentData((prev) => ({
                                  ...prev,
                                  theoryFile: e.target.files[0],
                                }))
                              }
                            />
                            <Button
                              className="mt-2 bg-green-600 hover:bg-green-700"
                              onClick={handleSaveAssignmentEdit}
                            >
                              Сохранить задание
                            </Button>
                          </>
                        ) : (
                          <>
                            <h4 className="font-semibold">{assignment.title}</h4>
                            {assignment.theoryFile && (
                              <p className="text-sm text-gray-600 mt-2">
                                Теория: {assignment.theoryFile.name}
                              </p>
                            )}
                            <Button
                              className="mt-2 bg-yellow-600 hover:bg-yellow-700"
                              onClick={() => handleEditAssignment(topicIndex, idx)}
                            >
                              Редактировать задание
                            </Button>
                          </>
                        )}
                      </div>
                    ))}

                    <div className="mt-4">
                      <Label htmlFor={`assignmentTitle-${topicIndex}`} className="block">
                        Название задания
                      </Label>
                      <TextInput
                        id={`assignmentTitle-${topicIndex}`}
                        type="text"
                        placeholder="Введите название задания"
                        value={(newAssignments[topicIndex] || {}).title || ""}
                        onChange={(e) =>
                          setNewAssignments((prev) => ({
                            ...prev,
                            [topicIndex]: {
                              ...prev[topicIndex],
                              title: e.target.value,
                            },
                          }))
                        }
                      />
                      <FileInput
                        id={`theoryFile-${topicIndex}`}
                        onChange={(e) =>
                          setNewAssignments((prev) => ({
                            ...prev,
                            [topicIndex]: {
                              ...prev[topicIndex],
                              theoryFile: e.target.files[0],
                            },
                          }))
                        }
                      />
                      <Button
                        className="mt-2 bg-blue-600 hover:bg-blue-700"
                        onClick={() => handleAddAssignment(topicIndex)}
                      >
                        Добавить задание
                      </Button>
                    </div>
                  </Accordion.Content>
                </Accordion.Panel>
              </Accordion>
            ))}

            <div className="mt-6">
              <Label htmlFor="newTopic" className="mb-2 block">Название темы</Label>
              <TextInput
                id="newTopic"
                type="text"
                placeholder="Введите название темы"
                value={newTopic}
                onChange={(e) => setNewTopic(e.target.value)}
              />
              <Button
                className="mt-4 bg-green-600 hover:bg-green-700"
                onClick={handleAddTopic}
              >
                Добавить тему
              </Button>
            </div>
          </div>

          <div className="flex justify-end">
            <Button className="bg-indigo-600 hover:bg-indigo-700" type="submit">
              Создать курс
            </Button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default CreateCoursePage;
