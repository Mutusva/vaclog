import React, { useState, useEffect, ChangeEvent } from "react";
import VacLogService from "../services/api/vaclogService";
import IRecordList from "../types/LogRecordList"
import IRecord from "../types/LogRecord";

const Records:React.FC =()=> {
  const [records, setRecords] = useState<Array<IRecord>>([]);
  const [currentRecord, setCurrentRecord] = useState<IRecord | null>(null);
  const [currentIndex, setCurrentIndex] = useState<number>(-1);
  const [searchDocumentID, setSearchDocumentID] = useState<string>("");

  useEffect(() => {
    getAllRecords();
  }, []);

  const onChangeSearchDocumentID = (e: ChangeEvent<HTMLInputElement>) => {
    const searchTitle = e.target.value;
    setSearchDocumentID(searchTitle);
  };

  const getAllRecords = () => {
    VacLogService.GetAllRecords()
      .then((response: any) => {
        setRecords(response.data.records);
        console.log(response.data.records);
      })
      .catch((e: Error) => {
        console.log(e);
      });
  };
  
  const refreshList = () => {
    getAllRecords();
    setCurrentRecord(null);
    setCurrentIndex(-1);
  };

  const setActiveRecord = (record: IRecord, index: number) => {
    setCurrentRecord(record);
    setCurrentIndex(index);
  };

  const findByDocumentID = () => {
    VacLogService.GetRecord(searchDocumentID)
      .then((response: any) => {
        setRecords(response.data);
        setCurrentRecord(null);
        console.log(response.data);
      })
      .catch((e: Error) => {
        console.log(e);
      });
  };


    return(
        <div className="list row">
      <div className="col-md-8">
        <div className="input-group mb-3">
          <input
            type="text"
            className="form-control"
            placeholder="Search by title"
            value={searchDocumentID}
            onChange={onChangeSearchDocumentID}
          />
          <div className="input-group-append">
            <button
              className="btn btn-outline-secondary"
              type="button"
              onClick={findByDocumentID}
            >
              Search
            </button>
          </div>
        </div>
      </div>
      <div className="col-md-6">
        <h4>Vaccination Records</h4>

        <ul className="list-group">
          {records &&
            records.map((record, index) => (
              <li
                className={
                  "list-group-item " + (index === currentIndex ? "active" : "")
                }
                onClick={() => setActiveRecord(record, index)}
                key={index}
              >
                {record.animal_id}
              </li>
            ))}
        </ul>
      </div>
      <div className="col-md-6">
        {currentRecord ? (
          <div>
            <h4>Record</h4>
            <div>
              <label>
                <strong>AnimalID:</strong>
              </label>{" "}
              {currentRecord.animal_id}
            </div>
            <div>
              <label>
                <strong>RecordId:</strong>
              </label>{" "}
              {currentRecord.document_id}
            </div>
            <div>
              <label>
                <strong>Date Administered:</strong>
              </label>{" "}
              {currentRecord.date_administered}
            </div>
            <div>
              <label>
                <strong>Species:</strong>
              </label>{" "}
              {currentRecord.species}
            </div>
            <div>
              <label>
                <strong>Age:</strong>
              </label>{" "}
              {currentRecord.age}
            </div>
            <div>
              <label>
                <strong>Vaccine Name:</strong>
              </label>{" "}
              {currentRecord.vaccine_name}
            </div>
            <div>
              <label>
                <strong>Notes:</strong>
              </label>{" "}
              {currentRecord.notes}
            </div>

            {/* <Link
              to={"/tutorials/" + currentTutorial.id}
              className="badge badge-warning"
            >
              Edit
            </Link> */}
          </div>
        ) : (
          <div>
            <br />
            <p>Please click on a Record...</p>
          </div>
        )}
      </div>
    </div>
    )
};

export default Records;