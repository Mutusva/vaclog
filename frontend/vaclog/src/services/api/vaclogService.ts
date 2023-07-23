import http from "../../http-common";
import IRecord from "../../types/LogRecord"
import IRecordList from "../../types/LogRecordList";


const GetAllRecords =() => {
    return http.get<IRecordList>("/records");
};

const CreateRecord =(data: IRecord) =>  {
    return http.put<any>("/records", data).
    catch((e)=>{
        console.log(e)
    });
};

const GetRecord =(docID: string) => {
    return http.get<IRecordList>(`/records/${docID}`);
};

const VacLogService = {
    GetAllRecords,
    CreateRecord,
    GetRecord
}

export default VacLogService;