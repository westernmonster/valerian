SELECT
    a.id,
    a.account_id,
    a.status,
    a.audit_conclusions,
    a.name,
    a.identification_number,
    a.id_card_type,
    a.id_card_start_date,
    a.id_card_expiry,
    a.address,
    a.sex,
    a.id_card_front_pic,
    a.id_card_back_pic,
    a.face_pic,
    a.ethnic_group,
    a.deleted,
    a.created_at,
    a.updated_at
FROM id_certifications a
WHERE a.id=:id AND a.deleted=0