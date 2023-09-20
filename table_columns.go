package models

type cols string

const (
	AccountsFeild cols = "account_id,company_unique_id,created_by,company_name,company_suffix,contact_info," +
		"time_zone,email,website,mc_number,cd_uname,cd_upasswd,cd_uid,cd_cookies,cd_search_cookies,cd_token_b_t," +
		"company_logo,company_description,bcc_address,departments_info,date_created,account_status,subscription_info," +
		"api_key,sandbox_api_key,settings_version,call_and_sms"
	DefaultSettingsFeild cols = "settings_id,company_id,new_lead_num,new_quote_num,new_order_num,lead_priority,required_deposit," +
		"required_type,first_qt_flw_up,quote_expire,asumd_delivered,ass_unvrifd_orders,remove_listing,issue_days,issue_mode,log_out," +
		"carrier_paymt_trm,carrier_paymt_begin,carrier_paymt_method,order_terms,dispatch_terms,quote_details_info,endpoints,deal_tags," +
		"payment_markup,request_for,payment_option,request_type,auto_lock,sendgrid_api_key,thank_you,tracking_script,hotjar_script," +
		"livechat_script,auto_quote,auto_quote_formula,notify_before_pickup,driver_before_temp,shipper_before_temp,notify_on_pickup," +
		"driver_on_temp,shipper_on_temp,system_sent_email,auto_responder_sett,auto_responder_body,return_url_pg_header," +
		"return_url_pg_subheader,return_url_pg_body,default_mailer"
	EmailsTemplatesField cols = "email_unique_id,email_id,company_id,email_type,name,used_for,for_follow_up,text_contents,html_contents," +
		"subject,description,attachment,cc_addresses"

	SMSTemplatesField cols = "sms_unique_id,sms_id,company_id,name,used_for,sms_body,description"

	AgentsField cols = "agent_id,company_id,name,email,phone,lead_multiples,assign_next_lead,leads_assigned,status,agent_role,super_active," +
		"super_agent_lead_no,username,password,level,last_history_id,last_email_check,agent_permissions,deal_list_settings," +
		"notifications_settings,commission_settings,webmail_settings,call_and_sms,drip_round_robins,dp_thumb,dp_large,date_joined," +
		"last_login,last_activity,total_reviews,ratings,total_recommendation,positive_feedback,cancellation_rate"

	AccntGrpField cols = "group_id,group_name,company_id,permissions"

	AgentsAccessField cols = "process_leads,process_quotes,process_orders,view_others_leads,process_others_leads,view_others_quotes," +
		"process_others_quotes,view_others_orders,process_others_orders,create_quotes,create_orders,crt_del_clients,crt_del_agents," +
		"crt_del_groups,crt_del_lead_src,view_stats,crt_edt_del_emails,crt_edt_del_forms,up_info_sett,generate_reports"

	Ungrouped_emails_templatesField cols = "email_id,company_id,to_email,subject,body,user_id,status,response_type,opens,clicks,date," +
		"tracking_id,email_unique_id,first_open"

	Quote_emails_templatesField cols = "email_id,quote_id,company_id,user_id,from_email,to_email,subject,body,cc_addresses,bcc_addresses," +
		"email_type,status,opens,clicks,date,tracking_id,email_unique_id,automation_uid,first_open,error_message"

	Order_emails_templatesField cols = "email_id,order_id,company_id,from_email,to_email,subject,body,cc_addresses,bcc_addresses,user_id," +
		"status,opens,clicks,date,tracking_id,email_unique_id,automation_uid,first_open,error_message"

	Batch_emailField cols = "batch_id,company_id,company_name,time_zone,agent_id,agent_name,agent_email,type,body,subject,cron_tab," +
		"integration_type,sendgrid_api,google_token,status,recipients,sent,opens,clicks,invalid_emails,unsubscribes,date,api_response"

	Followup_email_queueField cols = "queue_id,batch_id,company_id,user_id,quote_id,quote_uniq_id,from_email,from_name,to_email,to_name," +
		"next_followup,date,followup_eml,opens,first_open,status,tracking_id"

	CompanyAccountsField cols = "comp_account_id,comp_accnt_unique_id,company_id,account_type,username,firstname,lastname,fullname,email," +
		"company_name,tax_id,contacts_info,insurance_info,mc_number,cd_carrier_id,cd_carrier_sub_id,commission_structure,commission_amount," +
		"status,recent_block_date,recent_block_message,recent_block_by,date_added"

	MessagesField cols = "message_id,agent_id,company_id,preview,contact_email,contact_phone,contact_name,contact_type,has_thread,status," +
		"new_messages,last_unique_id,last_message_date"

	MessagesThreadField cols = "thread_id,message_id,unique_id,message_type,agent_id,deal_type,deal_id,company_id,contact_name," +
		"forwarding_email,send_from,send_to,subject,message_body,date,sent_by,status,outgoing_mail_status,outgoing_mail_info"

	CallHistoryField cols = "call_history_id,agent_id,company_id,RecordingSid,RecordingUrl,RecordingStatus,CallSid,RecordingStartTime," +
		"AccountSid,RecordingDuration,FromNumbr,ToNumbr,item,item_id,date"

	Pickup_dropoff_contactsField cols = "pickdrop_account_id,company_id,account_id,firstname,lastname,fullname,email,company_name," +
		"contacts_info,status,account_notes,date_added"

	Carrier_driversField cols = "driver_id,company_id,account_id,firstname,lastname,email,phone_1,status,date_added"

	Lead_sourceField cols = "source_id,company_id,source_email,api_access_key,price_per_lead,status,distributed,cron_processed,company_info," +
		"auto_quote,autoqouote_formular,use_for_extnl_form,deflt_price_checker,automation_id,website_url,last_email_deleted,got_emails," +
		"settings,import_mapping"

	Lead_source_settField cols = "settings_id,leads_source_id,company_id,agent_id,lead_multiples,assign_next_lead,leads_assigned,route_leads," +
		"auto_quote,auto_quote_formula,use_for_extnl_form"

	LeadsField cols = "lead_id,lead_uniq_id,company_id,dates,duplicate_info,car_run,ship_via,assigned_to,assigned_to_info,original_assigned_to," +
		"referral,shipper,origin,destination,vehicles,source,source_type,source_id,payments,type,seen_by_agent,original_email,item_counts,tags," +
		"lead_token,received,referral_id"

	QuotesField cols = "assigned_to,assigned_to_info,automate,automation_id,automation_status,car_run,company_id,dates,destination," +
		"duplicate_info,follow_up_info,item_counts,next_followup_date,origin,original_assigned_to,original_email,payments,quote_id,quote_token," +
		"quote_uniq_id,quoted,referral,referral_id,seen_by_agent,ship_via,shipper,shpr_shipper_id,source,source_type,source_id,tags,type," +
		"vehicles"

	OrdersField cols = "order_id,order_uniq_id,company_id,dates,mode,car_run,ship_via,assigned_to,assigned_to_info,original_assigned_to," +
		"source,source_type,source_id,referral,payments,shipper,origin,destination,vehicles,load_on,deliver_on,posted_board,hdesk_post_id," +
		"dispatch_details,status,type,resolve_issue_to,original_email,quote_id_converted,quote_uniq_id_conv,tracking_number,converted_details," +
		"order_token,cd_dispatch_id,cd_dispatch_status,user_order_id,automate,driver_id,carrier_id,shipper_signature,driver_signature," +
		"subscribed_drips,automation_id,automation_status,tags,item_counts,issues,dates_created,shpr_shipper_id,dispatched_to,p_terminal_id," +
		"d_terminal_id,referral_id"

	ImportedDealFilesField cols = "can_start_import,data_headers,data_rows,data_type,date_added,file_location,import_id,imported," +
		"imported_data,rejected_leads,total_file_data,uploaded_by"
)
