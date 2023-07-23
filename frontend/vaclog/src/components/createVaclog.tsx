import React, { useState, ChangeEvent } from "react";
import VacLogService from "../services/api/vaclogService";
import IRecord from '../types/LogRecord';


const CreateRecord: React.FC = () => {

    const initialRecordState = {
        animal_id: "",
        species: "",
        age: 0,
        vaccine_name: "",
        date_administered: "",
        notes: ""
    };

    const initailCreateResponse = {
        transactionId: "",
        documentId: ""
    }

    const [record, setRecord] = useState<IRecord>(initialRecordState);
    const [createRecordResponse, setCreateRecordResponse]= useState<any>(initailCreateResponse);
    const [submitted, setSubmitted] = useState<boolean>(false);
    const [errors, setErrors] = useState([]);
    const [hasErrors, setHasErrors] = useState(false);

    const handleInputChange = (event: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setRecord({ ...record, [name]: value });
    };

    const saveRecord = () => {
        var data = {
            animal_id: record.animal_id,
            species: record.species,
            age: record.age,
            vaccine_name: record.vaccine_name,
            date_administered: record.date_administered,
            notes: record.notes
        };
    
        VacLogService.CreateRecord(data)
          .then((response: any) => {
            setCreateRecordResponse({
              transactionId: response.data.transactionId,
              documentId: response.data.documentId,
            });
            setSubmitted(true);
            console.log(response.data);
          })
          .catch((e: Error) => {
            console.log(e);
            setHasErrors(true)

          });
      };

      const newRecord = () => {
        setRecord(initialRecordState);
        setSubmitted(false);
        setHasErrors(false);
      };

    return (
        <div className="submit-form">
        
        {
            hasErrors && (
                <div className="alert alert-danger" role="alert">
                    Error creating record
                </div>
            )
        }
        
        {submitted ? (
          <div>
            <h4>You submitted successfully!</h4>
            <button className="btn btn-success" onClick={newRecord}>
              Add Record
            </button>
          </div>
        ) : (
          <div>
            <div className="form-group">
              <label htmlFor="animal_id">Animal ID</label>
              <input
                type="text"
                className="form-control"
                id="animal_id"
                required={true}
                value={record.animal_id}
                onChange={handleInputChange}
                name="animal_id"
              />
            </div>
  
            <div className="form-group">
              <label htmlFor="species">Species</label>
              <input
                type="text"
                className="form-control"
                id="species"
                required={true}
                value={record.species}
                onChange={handleInputChange}
                name="species"
              />
            </div>

            <div className="form-group">
              <label htmlFor="age">Age (In Months)</label>
              <input
                type="number"
                className="form-control"
                id="age"
                required={true}
                value={record.age}
                onChange={handleInputChange}
                name="age"
              />
            </div>

            <div className="form-group">
              <label htmlFor="vaccine_name">Vaccine Name</label>
              <input
                type="text"
                className="form-control"
                id="vaccine_name"
                required={true}
                value={record.vaccine_name}
                onChange={handleInputChange}
                name="vaccine_name"
              />
            </div>

            <div className="form-group">
              <label htmlFor="date_administered">Date Administered</label>
              <input
                type="date"
                className="form-control"
                id="date_administered"
                required={true}
                value={record.date_administered}
                onChange={handleInputChange}
                name="date_administered"
              />
            </div>

            <div className="form-group">
              <label htmlFor="notes">Notes</label>
              <input
                type="text"
                className="form-control"
                id="notes"
                required={true}
                value={record.notes}
                onChange={handleInputChange}
                name="notes"
              />
            </div>
  
            <button onClick={saveRecord} className="btn btn-success">
              Submit
            </button>
          </div>
        )}
      </div>
    )
}

export default CreateRecord;