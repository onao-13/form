const API = "http://37.140.241.38:8150//api/send";

document.querySelector("#send").addEventListener('click', () => send());

async function send() {
    var organizationName = document.querySelector("#organization_name").value;
    var shortOrganizationName = document.querySelector("#short_organization_name").value;
    var typeEmployment = document.querySelector("#type_employment").value;
    var positionLeaderOrganization = document.querySelector("#position_leader_organization").value;
    var fioLeaderOrganization = document.querySelector("#fio_leader_organization").value;
    var positionResponsiblePerson = document.querySelector("#position_responsible_person").value;
    var fioResponsiblePerson = document.querySelector("#fio_responsible_person").value;
    var number = document.querySelector("#number").value;
    var email = document.querySelector("#email").value;

    const body = JSON.stringify({
        "organization_name": organizationName,
        "short_organization_name": shortOrganizationName,
        "type_participant": typeEmployment,
        "position_leader_organization": positionLeaderOrganization,
        "fio_leader_organization": fioLeaderOrganization,
        "position_responsible_person": positionResponsiblePerson,
        "fio_responsible_person": fioResponsiblePerson,
        "telephone_number": number,
        "email": email
    })

    var res = await fetch(API, {
        method: "POST",
        body: body,
    })

    var titleModal = document.querySelector("#modal-title");
    var errorModal = document.querySelector("#error");

    if (!res.ok) {
        titleModal.innerHTML = "Ошибка отправки заявки";

        var res = await res.json();
        errorModal.innerHTML = res.error;
    } else {
        titleModal.innerHTML = "Данные отправлены";
        errorModal.innerHTML = "Данные успешно отправлены нам! Скоро мы с вами свяжемся";
    }
}
